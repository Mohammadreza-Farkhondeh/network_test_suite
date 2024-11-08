import time
import subprocess
import csv
import signal
from datetime import datetime
from threading import Event
import speedtest
from ping3 import ping
import plot_network_results as plot

stop_event = Event()


def handle_exit_signal(signum, frame):
    print("Exit signal received. Stopping the monitoring...")
    stop_event.set()


signal.signal(signal.SIGINT, handle_exit_signal)

FILE_NAME = f"results/network_status_{datetime.now()}.csv"
INTERVAL = 60
DURATION = 86400
services_to_check = ["net2.sharif.edu", "sharif.edu", "ictc.sharif.edu"]
dns_servers = ["172.26.146.34", "172.26.146.35"]


def run_ping(target, count=5):
    response = subprocess.run(
        ["ping", f"-c {count}", target],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        text=True,
    )
    return response


def parse_ping_response(response):
    if response.returncode == 0:
        output_lines = response.stdout.split("\n")
        latencies = [
            float(line.split("time=")[1].split(" ")[0])
            for line in output_lines
            if "time=" in line
        ]
        return latencies
    return None


def ping_target(target, count=5):
    results = []
    for _ in range(count):
        response = ping(target, timeout=2)
        if response is not None:
            results.append(response * 1000)
    return results


def calculate_latency(latencies):
    return sum(latencies) / len(latencies) if latencies else None


def calculate_jitter(latencies):
    return max(latencies) - min(latencies) if latencies and len(latencies) > 1 else None


def calculate_packet_loss(latencies, count):
    return 100 - (len(latencies) / count * 100) if latencies else 100


def get_latency_jitter_packet_loss(target="8.8.8.8", count=5):
    latencies = ping_target(target, count)
    latency = calculate_latency(latencies)
    jitter = calculate_jitter(latencies)
    packet_loss = calculate_packet_loss(latencies, count)
    return latency, jitter, packet_loss


def run_speedtest():
    st = speedtest.Speedtest()
    st.download()
    st.upload()
    return st


def get_download_speed(st):
    return st.results.download / 1_000_000


def get_upload_speed(st):
    return st.results.upload / 1_000_000


def get_bandwidth():
    try:
        st = run_speedtest()
        download = get_download_speed(st)
        upload = get_upload_speed(st)
    except Exception as ex:
        print("errrr in speedtest" , ex)
        download, upload = None, None
    return download, upload


def ping_single_target(target):
    response = subprocess.run(
        ["ping", "-c", "1", target],
        stdout=subprocess.PIPE,
        stderr=subprocess.PIPE,
        text=True,
    )
    return response.returncode == 0


def check_accessibility(targets):
    results = {}
    for target in targets:
        accessible = ping_single_target(target)
        results[target] = accessible
    return results


def gather_network_metrics():
    latency, jitter, packet_loss = get_latency_jitter_packet_loss()
    download, upload = get_bandwidth()
    return latency, jitter, packet_loss, download, upload


def gather_accessibility_metrics():
    return check_accessibility(services_to_check + dns_servers)


def open_csv_file(filename):
    file_exists = False
    try:
        with open(filename, "r") as f:
            file_exists = True
    except FileNotFoundError:
        pass
    return file_exists


def prepare_csv_writer(filename, fieldnames):
    file_exists = open_csv_file(filename)
    csvfile = open(filename, "a", newline="")
    writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
    if not file_exists:
        writer.writeheader()
    return writer, csvfile


def close_csv_file(csvfile):
    csvfile.close()


def get_timestamp():
    return datetime.now().strftime("%Y-%m-%d %H:%M:%S")


def prepare_log_data():
    network_metrics = gather_network_metrics()
    accessibility_metrics = gather_accessibility_metrics()
    return {
        "timestamp": get_timestamp(),
        "latency_ms": network_metrics[0],
        "jitter_ms": network_metrics[1],
        "packet_loss_percent": network_metrics[2],
        "bandwidth_download_mbps": network_metrics[3],
        "bandwidth_upload_mbps": network_metrics[4],
        **{
            f"access_{service}": int(accessible)
            for service, accessible in accessibility_metrics.items()
        },
    }


def log_network_status(filename="network_status.csv"):
    fieldnames = [
        "timestamp",
        "latency_ms",
        "jitter_ms",
        "packet_loss_percent",
        "bandwidth_download_mbps",
        "bandwidth_upload_mbps",
        *[f"access_{service}" for service in services_to_check + dns_servers],
    ]
    writer, csvfile = prepare_csv_writer(filename, fieldnames)
    log_data = prepare_log_data()
    writer.writerow(log_data)
    close_csv_file(csvfile)


def start_network_monitoring(duration=86400, interval=600, filename=FILE_NAME):
    end_time = time.time() + duration
    while time.time() < end_time and not stop_event.is_set():
        log_network_status(filename)
        for _ in range(int(interval / 10)):
            if stop_event.is_set():
                break
            time.sleep(10)
    print("Testing finished. Plotting results...")
    plot.main(file_name=FILE_NAME)


if __name__ == "__main__":
    start_network_monitoring(duration=DURATION, interval=INTERVAL)
