.PHONY: test all build clean run check cover lint docker help
BIN_FILE=Appointment
all: check build
build:
	@go build -o "${BIN_FILE}"
clean:
	@go clean
	rm --force "xx.out"
clean_all: clean
	rm --force database.db
test:
	@go test
check:
	@go fmt ./
	@go vet ./
cover:
	@go test -coverprofile xx.out
	@go tool cover -html=xx.out
./"${BIN_FILE}": build
run: ./"${BIN_FILE}"
	./"${BIN_FILE}"
lint:
	golangci-lint run --enable-all
help:
	@echo "make 格式化go代码 并编译生成二进制文件"
	@echo "make build 编译go代码生成二进制文件"
	@echo "make clean 清理中间目标文件"
	@echo "make clean_all 清理中间目标文件以及database.db"
	@echo "make test 执行测试case"
	@echo "make check 格式化go代码"
	@echo "make cover 检查测试覆盖率"
	@echo "make run 直接运行程序"
	@echo "make lint 执行代码检查"
