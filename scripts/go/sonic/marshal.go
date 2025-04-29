package main

import (
	"encoding/json"
	"fmt"

	"github.com/bytedance/sonic"
)

func main() {
	// 定义一个包含特殊字符（换行符）的结构体
	val1 := "\n{\"__write_time\":\"2025-04-25T08:36:33.587Z\",\"category\":\"metric\",\"__index_base\":\"emma\",\"tags\":\"warm\", \"@version\":\"1\",\"__routing\":\"954868\",\n\"labels\":{\"host_name\":\"node-15-02\",\"host_ip\":\"10.4.15.02\",\n\"cmdline\":\"/app/deploy-service/deploy-service\",\"job_id\":\"job-3f0b416037baaafb\",\"process_id\":\"33751\",\"memory_stat\":\"RSS\",\"process_name\":\"deploy-service\"}\n,\"@timestamp\":\"2025-04-25T08:36:33.587Z\",\"type\":\"sys\",\"metrics\":{\"node_processes_memory_bytes_total\":2036941271004109105},\"__data_type\":\"emma\"}"
	root, _ := sonic.Get([]byte(val1))
	raw, _ := root.Raw()
	fmt.Println(raw)

	fmt.Println("-----------------------")
	buf, _ := root.MarshalJSON()
	fmt.Println(string(buf))

	fmt.Println("-----------------------")
	exp, _ := json.Marshal(&root)
	fmt.Println(string(exp))
}
