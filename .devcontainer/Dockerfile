FROM archlinux:latest

RUN pacman-key --init
RUN pacman -Syyu --noconfirm

RUN pacman -S --noconfirm \
    zsh sudo git vim \
    bat zip unzip tar wget \
    curl fd htop ncdu mlocate \
    tree man-db tldr jq yq \
    diffutils openssl go base-devel

RUN useradd -m vscode \
    && usermod -aG wheel vscode \
    && echo '%wheel ALL=(ALL) NOPASSWD:ALL' | tee /etc/sudoers.d/wheel > /dev/null

USER vscode
WORKDIR /home/vscode

RUN git clone https://aur.archlinux.org/yay-bin.git && \
    cd yay-bin && makepkg -si --noconfirm && cd .. && \
    rm -rf ./yay-bin

RUN yay -S --noconfirm pfetch-rs-bin

RUN sh -c '$(curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh)' && \
    git clone https://github.com/zsh-users/zsh-syntax-highlighting.git ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/zsh-syntax-highlighting && \
    git clone https://github.com/zdharma-continuum/fast-syntax-highlighting ${ZSH_CUSTOM:-~/.oh-my-zsh/custom}/plugins/fast-syntax-highlighting && \
    sed -i 's/^plugins=(*)/plugins=(git ssh zsh-syntax-highlighting)/g' ~/.zshrc

RUN curl -s https://ohmyposh.dev/install.sh | bash -s && \
    echo -e '\neval $(oh-my-posh init zsh)' >> ~/.zshrc

CMD pfetch && zsh
