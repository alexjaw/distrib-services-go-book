FROM gitpod/workspace-full

RUN sudo apt-get update \
 && sudo apt-get install -y \
 protobuf-compiler \
 golang-goprotobuf-dev \
 && sudo rm -rf /var/lib/apt/lists/*