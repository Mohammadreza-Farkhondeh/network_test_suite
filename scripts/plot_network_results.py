import pandas as pd
import matplotlib.pyplot as plt
from matplotlib.dates import DateFormatter


def load_data(filename):
    df = pd.read_csv(filename)
    df["timestamp"] = pd.to_datetime(df["timestamp"])
    return df


def plot_latency(df, filename="latency_plot.png"):
    plt.figure(figsize=(20, 9))
    plt.plot(
        df["timestamp"], df["latency_ms"], marker="o", color="b", label="Latency (ms)"
    )
    plt.title("Network Latency Over Time")
    plt.xlabel("Time")
    plt.ylabel("Latency (ms)")
    plt.legend()
    plt.grid(True)
    plt.gca().xaxis.set_major_formatter(DateFormatter("%H:%M"))
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


def plot_jitter(df, filename="jitter_plot.png"):
    plt.figure(figsize=(20, 9))
    plt.plot(
        df["timestamp"],
        df["jitter_ms"],
        marker="o",
        color="orange",
        label="Jitter (ms)",
    )
    plt.title("Network Jitter Over Time")
    plt.xlabel("Time")
    plt.ylabel("Jitter (ms)")
    plt.legend()
    plt.grid(True)
    plt.gca().xaxis.set_major_formatter(DateFormatter("%H:%M"))
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


def plot_packet_loss(df, filename="packet_loss_plot.png"):
    plt.figure(figsize=(20, 9))
    plt.plot(
        df["timestamp"],
        df["packet_loss_percent"],
        marker="o",
        color="red",
        label="Packet Loss (%)",
    )
    plt.title("Network Packet Loss Over Time")
    plt.xlabel("Time")
    plt.ylabel("Packet Loss (%)")
    plt.legend()
    plt.grid(True)
    plt.gca().xaxis.set_major_formatter(DateFormatter("%H:%M"))
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


def plot_bandwidth(df, filename="bandwidth_plot.png"):
    plt.figure(figsize=(20, 9))
    plt.plot(
        df["timestamp"],
        df["bandwidth_download_mbps"],
        marker="o",
        color="green",
        label="Download Bandwidth (Mbps)",
    )
    plt.plot(
        df["timestamp"],
        df["bandwidth_upload_mbps"],
        marker="o",
        color="purple",
        label="Upload Bandwidth (Mbps)",
    )
    plt.title("Network Bandwidth Over Time")
    plt.xlabel("Time")
    plt.ylabel("Bandwidth (Mbps)")
    plt.legend()
    plt.grid(True)
    plt.gca().xaxis.set_major_formatter(DateFormatter("%H:%M"))
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


def plot_accessibility(df, filename="accessibility_plot.png"):
    plt.figure(figsize=(20, 9))
    accessibility_columns = [col for col in df.columns if col.startswith("access_")]
    for col in accessibility_columns:
        plt.plot(df["timestamp"], df[col], marker="o", label=col)
    plt.title("Service Accessibility Over Time")
    plt.xlabel("Time")
    plt.ylabel("Accessibility (1 = accessible, 0 = not accessible)")
    plt.legend()
    plt.grid(True)
    plt.gca().xaxis.set_major_formatter(DateFormatter("%H:%M"))
    plt.xticks(rotation=45)
    plt.tight_layout()
    plt.savefig(filename)
    plt.close()


def main(*args, **kwargs):
    file_name = kwargs.get("file_name")
    df = load_data(file_name)

    plot_latency(df, "results/latency_plot.png")
    plot_jitter(df, "results/jitter_plot.png")
    plot_packet_loss(df, "results/packet_loss_plot.png")
    plot_bandwidth(df, "results/bandwidth_plot.png")
    plot_accessibility(df, "results/accessibility_plot.png")
