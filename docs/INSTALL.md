## Install
* Mac OS (m1)
  ```shell
  curl -s https://api.github.com/repos/evgeny-klyopov/blockchain-node-export/releases/latest \
  | grep browser_download_url \
  | grep bne-macos-m1.tar.gz \
  | cut -d '"' -f 4 \
  | wget -qi - 
  tar -xvf bne-macos-m1.tar.gz && mv bne /usr/bin/bne 
  ```
* Mac OS
  ```shell
  curl -s https://api.github.com/repos/evgeny-klyopov/blockchain-node-export/releases/latest \
  | grep browser_download_url \
  | grep bne-macos.tar.gz \
  | cut -d '"' -f 4 \
  | wget -qi - 
  tar -xvf bne-macos.tar.gz && mv bne /usr/bin/bne 
  ```
* Linux
  ```shell
  curl -s https://api.github.com/repos/evgeny-klyopov/blockchain-node-export/releases/latest \
  | grep browser_download_url \
  | grep bne-linux.tar.gz \
  | cut -d '"' -f 4 \
  | wget -qi - 
  tar -xvf bne-linux.tar.gz && mv bne /usr/bin/bne 
  ```
* Windows
  ```shell
  curl -s https://api.github.com/repos/evgeny-klyopov/blockchain-node-export/releases/latest \
  | grep browser_download_url \
  | grep bne-windows.tar.gz \
  | cut -d '"' -f 4 \
  | wget -qi -
  tar -xvf bne-windows.tar.gz && mv bne.exe /usr/bin/bne.exe 
  ```