#
# NOTE: THIS DOCKERFILE IS GENERATED VIA "apply-templates.sh"
#
# PLEASE DO NOT EDIT IT DIRECTLY.
#

FROM buildpack-deps:bookworm-scm

# install cgo-related dependencies
RUN set -eux; \
	apt-get update; \
	apt-get install -y --no-install-recommends \
		g++ \
		gcc \
		libc6-dev \
		make \
		pkg-config \
	; \
	rm -rf /var/lib/apt/lists/*

ENV PATH /usr/local/go/bin:$PATH

ENV GOLANG_VERSION 1.21.4

RUN set -eux; \
	arch="$(dpkg --print-architecture)"; arch="${arch##*-}"; \
	url=; \
	case "$arch" in \
		'amd64') \
			url='https://dl.google.com/go/go1.21.4.linux-amd64.tar.gz'; \
			sha256='73cac0215254d0c7d1241fa40837851f3b9a8a742d0b54714cbdfb3feaf8f0af'; \
			;; \
		'armel') \
			export GOARCH='arm' GOARM='5' GOOS='linux'; \
			;; \
		'armhf') \
			url='https://dl.google.com/go/go1.21.4.linux-armv6l.tar.gz'; \
			sha256='6c62e89113750cc77c498194d13a03fadfda22bd2c7d44e8a826fd354db60252'; \
			;; \
		'arm64') \
			url='https://dl.google.com/go/go1.21.4.linux-arm64.tar.gz'; \
			sha256='ce1983a7289856c3a918e1fd26d41e072cc39f928adfb11ba1896440849b95da'; \
			;; \
		'i386') \
			url='https://dl.google.com/go/go1.21.4.linux-386.tar.gz'; \
			sha256='64d3e5d295806e137c9e39d1e1f10b00a30fcd5c2f230d72b3298f579bb3c89a'; \
			;; \
		'mips64el') \
			url='https://dl.google.com/go/go1.21.4.linux-mips64le.tar.gz'; \
			sha256='c7ce3a9dcf03322b79beda474c4a0154393d9029b48f7c2e260fb3365c8a6ad3'; \
			;; \
		'ppc64el') \
			url='https://dl.google.com/go/go1.21.4.linux-ppc64le.tar.gz'; \
			sha256='2c63b36d2adcfb22013102a2ee730f058ec2f93b9f27479793c80b2e3641783f'; \
			;; \
		'riscv64') \
			url='https://dl.google.com/go/go1.21.4.linux-riscv64.tar.gz'; \
			sha256='9695edd2109544b364daddb32816f5c7980f1f48b8490c51fa2c167f5b2eca48'; \
			;; \
		's390x') \
			url='https://dl.google.com/go/go1.21.4.linux-s390x.tar.gz'; \
			sha256='7a75ba4afc7a96058ca65903d994cd862381825d7dca12b2183f087c757c26c0'; \
			;; \
		*) echo >&2 "error: unsupported architecture '$arch' (likely packaging update needed)"; exit 1 ;; \
	esac; \
	build=; \
	if [ -z "$url" ]; then \
# https://github.com/golang/go/issues/38536#issuecomment-616897960
		build=1; \
		url='https://dl.google.com/go/go1.21.4.src.tar.gz'; \
		sha256='47b26a83d2b65a3c1c1bcace273b69bee49a7a7b5168a7604ded3d26a37bd787'; \
		echo >&2; \
		echo >&2 "warning: current architecture ($arch) does not have a compatible Go binary release; will be building from source"; \
		echo >&2; \
	fi; \
	\
	wget -O go.tgz.asc "$url.asc"; \
	wget -O go.tgz "$url" --progress=dot:giga; \
	echo "$sha256 *go.tgz" | sha256sum -c -; \
	\
# https://github.com/golang/go/issues/14739#issuecomment-324767697
	GNUPGHOME="$(mktemp -d)"; export GNUPGHOME; \
# https://www.google.com/linuxrepositories/
	gpg --batch --keyserver keyserver.ubuntu.com --recv-keys 'EB4C 1BFD 4F04 2F6D DDCC  EC91 7721 F63B D38B 4796'; \
# let's also fetch the specific subkey of that key explicitly that we expect "go.tgz.asc" to be signed by, just to make sure we definitely have it
	gpg --batch --keyserver keyserver.ubuntu.com --recv-keys '2F52 8D36 D67B 69ED F998  D857 78BD 6547 3CB3 BD13'; \
	gpg --batch --verify go.tgz.asc go.tgz; \
	gpgconf --kill all; \
	rm -rf "$GNUPGHOME" go.tgz.asc; \
	\
	tar -C /usr/local -xzf go.tgz; \
	rm go.tgz; \
	\
	if [ -n "$build" ]; then \
		savedAptMark="$(apt-mark showmanual)"; \
		apt-get update; \
		apt-get install -y --no-install-recommends golang-go; \
		\
		export GOCACHE='/tmp/gocache'; \
		\
		( \
			cd /usr/local/go/src; \
# set GOROOT_BOOTSTRAP + GOHOST* such that we can build Go successfully
			export GOROOT_BOOTSTRAP="$(go env GOROOT)" GOHOSTOS="$GOOS" GOHOSTARCH="$GOARCH"; \
			./make.bash; \
		); \
		\
		apt-mark auto '.*' > /dev/null; \
		apt-mark manual $savedAptMark > /dev/null; \
		apt-get purge -y --auto-remove -o APT::AutoRemove::RecommendsImportant=false; \
		rm -rf /var/lib/apt/lists/*; \
		\
# remove a few intermediate / bootstrapping files the official binary release tarballs do not contain
		rm -rf \
			/usr/local/go/pkg/*/cmd \
			/usr/local/go/pkg/bootstrap \
			/usr/local/go/pkg/obj \
			/usr/local/go/pkg/tool/*/api \
			/usr/local/go/pkg/tool/*/go_bootstrap \
			/usr/local/go/src/cmd/dist/dist \
			"$GOCACHE" \
		; \
	fi; \
	\
	go version

# don't auto-upgrade the gotoolchain
# https://github.com/docker-library/golang/issues/472
ENV GOTOOLCHAIN=local

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 1777 "$GOPATH"
WORKDIR $GOPATH

# Install Formatter 2023/11/24
RUN	go install golang.org/x/tools/cmd/goimports@latest
RUN	go install github.com/rakyll/hey@latest
RUN	go install honnef.co/go/tools/cmd/staticcheck@latest
RUN	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
RUN	go install golang.org/x/tools/gopls@latest