# Corel

**Corel** is a minimalist, autonomous lateral movement agent designed for high-performance discovery and self-replication within containerized environments. 

Unlike traditional static relays, **Corel** is built to "hop" between nodes, leveraging local environment primitives to replicate its core binary and establish a distributed, decentralized presence across **Kubernetes** and (**Linux** | **Windows** | **MacOs**) clusters.

---

## ⚡ Core Philosophy

Corel follows a strictly autonomous **Scan-Infect-Replicate** lifecycle:

1.  **Discovery**: Passive and active identification of adjacent nodes using internal network scans, `/etc/hosts` parsing, or K8s API queries.
2.  **Validation**: Assessment of target accessibility (SSH, exposed sockets, or vulnerable APIs).
3.  **Delivery**: The agent reads its own memory/binary and pushes it to the target host.
4.  **Execution**: Remote initialization and residency, turning the target into a new "Core" for further propagation.

## ✨ Key Features

* **Self-Replication**: Built-in logic to handle binary transfer via SFTP, Base64-streaming, or HTTP-pull.
* **Zero-Dependency**: Compiled as a fully static binaire (CGO disabled) for immediate execution on `alpine`, `scratch`, or `distroless` images.
* **K8s-Native**: Optional integration with `ServiceAccount` tokens to automate sidecar injection and pod-to-pod movement.
* **Modular Propagators**: Extensible interface system to add new movement vectors (SSH, Kube-API, Docker-Socket).
* **Stealth focused**: Minimal footprint, binary stripping, and low-latency scanning.

## ⚠️ Disclaimer

This project is **strictly for educational purposes** and security research only. It is designed to help understand lateral movement techniques and containerized environment security.

**Any illegal or malicious use is strictly prohibited.** The author(s) take no responsibility for any misuse of this tool. Always obtain proper authorization before testing on any system you do not own.


## 👤 Credits

Developped by **@lukas-sgx**.  
Thank's community for feedbacks.

## 📄 Licence

This project is distributed on [`LICENSE MIT`](./LICENSE)