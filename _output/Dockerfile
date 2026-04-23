FROM fedora:44
RUN dnf install -y hwdata pciutils && dnf clean all
COPY drasriov /drasriov
ENTRYPOINT ["/drasriov"]
