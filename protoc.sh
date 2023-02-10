proto_dirs=$(find ./proto -path -prune -o -name '*.proto' -print0 | xargs -0 -n1 dirname | sort | uniq)
for dir in $proto_dirs; do
    protoc \
    -I "proto" \
    -I "third_party" \
    --go_out=plugins=grpc:. \
    --grpc-gateway_out=logtostderr=true:. \
  $(find "${dir}" -maxdepth 1 -name '*.proto')

done



# protoc \
#  -I "third_party/" \
#  -I "proto" \
# --grpc-gateway_out=logtostderr=true:. \
#  --go_out=plugins=grpc:. \
#   ./proto/mathv2.proto