go tool dist list | tr '\/' ' ' | while read targetos targetarch
do
  echo "GOOS=${targetos} GOARCH=${targetarch} go build -o choose-items_${targetos}_${targetarch}"
  GOOS=${targetos} GOARCH=${targetarch} go build -o choose-items_${targetos}_${targetarch}
done

# EOF
# darwin amd64
# darwin arm64
# linux amd64
# linux arm
# linux arm64
# windows amd64
# windows arm64
# EOF
