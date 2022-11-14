package main

import (
	mailX "awesomeProject/mail"
	"awesomeProject/mode"
	"awesomeProject/resource"
	"awesomeProject/utils"
	"flag"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/xuri/excelize/v2"
	"strconv"
)

func init() {
	//读取配置文件
	var config string
	flag.StringVar(&config, "c", "./conf.toml", "默认路径默认即可")
	flag.Parse()
	_, err := toml.DecodeFile(config, &model.C)
	if err != nil {
		fmt.Println(err)
	}
}

func CreateXlS(up, cpu, mem, disk, netEs, netTW, DwRite, Dread, NetU, NetD []string, fileName string) {
	f := excelize.NewFile()
	var headerNameArray = []string{"主机状态(UP)", "CPU使用率%", "内存使用率%", "磁盘使用率%", "网络总线程数", "网络TW数", "磁盘IO写MB/s", "磁盘IO读MB/s", "网络Up-MB/s", "网络Down-MB/s"}
	sheetName := "sheet1"
	sheetWords := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U",
		"V", "W", "X", "Y", "Z",
	}

	for k, v := range headerNameArray {
		f.SetCellValue(sheetName, sheetWords[k]+"1", v)
	}

	//设置列的宽度
	f.SetColWidth("Sheet1", "A", sheetWords[len(headerNameArray)-1], 35)
	headStyleID, _ := f.NewStyle(`{ 
   "font":{
      "color":"#333333",
      "bold":false, 
      "family":"arial"
 
   }, 
   "alignment":{ 
      "vertical":"center", 
      "horizontal":"center" 
   } 
}`)
	//设置表头的样式
	f.SetCellStyle(sheetName, "A1", sheetWords[len(headerNameArray)-1]+"1", headStyleID)
	u := 1
	c := 1
	m := 1
	d := 1
	ne := 1
	ntw := 1
	dw := 1
	dr := 1
	nu := 1
	nd := 1
	for _, vp := range up {
		u++
		f.SetCellValue(sheetName, sheetWords[0]+strconv.Itoa(u), vp)
	}
	for _, vc := range cpu {
		c++
		f.SetCellValue(sheetName, sheetWords[1]+strconv.Itoa(c), vc)
	}
	for _, vm := range mem {
		m++
		f.SetCellValue(sheetName, sheetWords[2]+strconv.Itoa(m), vm)
	}
	for _, vd := range disk {
		d++
		f.SetCellValue(sheetName, sheetWords[3]+strconv.Itoa(d), vd)
	}
	for _, vne := range netEs {
		ne++
		f.SetCellValue(sheetName, sheetWords[4]+strconv.Itoa(ne), vne)
	}
	for _, vtw := range netTW {
		ntw++
		f.SetCellValue(sheetName, sheetWords[5]+strconv.Itoa(ntw), vtw)
	}
	for _, vdw := range DwRite {
		dw++
		f.SetCellValue(sheetName, sheetWords[6]+strconv.Itoa(dw), vdw)
	}
	for _, vdr := range Dread {
		dr++
		f.SetCellValue(sheetName, sheetWords[7]+strconv.Itoa(dr), vdr)
	}
	for _, vnu := range NetU {
		nu++
		f.SetCellValue(sheetName, sheetWords[8]+strconv.Itoa(nu), vnu)
	}
	for _, vnd := range NetD {
		nd++
		f.SetCellValue(sheetName, sheetWords[9]+strconv.Itoa(nd), vnd)
	}

	if err := f.SaveAs(fileName + ".xlsx"); err != nil {
		fmt.Println(err)
	}

	if _, err := f.WriteToBuffer(); err != nil {
		fmt.Println(err)
	}
}
func PasteData(jobs ...string) {
	for _, job := range jobs {
		up := resource.GetData("up", job)
		c := fmt.Sprintf(`(1-avg(rate(node_cpu_seconds_total{job="%s",mode='idle'}[30s]))by(instance))*100`, job)
		cpu := resource.GetCCData(c)
		m := fmt.Sprintf(`(1-(node_memory_MemAvailable_bytes{job="%s"}/(node_memory_MemTotal_bytes{job="%s"})))*100`, job, job)
		m6 := `(1-(node_memory_MemFree_bytes%7Bjob%3D"Centos6-host"%7D%2Bnode_memory_Buffers_bytes%7Bjob%3D"Centos6-host"%7D%2Bnode_memory_Cached_bytes%7Bjob%3D"Centos6-host"%7D)%2F(node_memory_MemTotal_bytes%7Bjob%3D"Centos6-host"%7D))*100`
		var mem []string
		if job == "host-node" {
			mem = resource.GetMData(m)

		} else {
			mem = resource.GetMData(m6)
		}
		var disk []string
		//d := fmt.Sprintf(`(node_filesystem_size_bytes{job="%s",mountpoint="/",fstype=~"ext.?|xfs"}-node_filesystem_free_bytes{job="%s",mountpoint="/",fstype=~"ext.?|xfs"})/node_filesystem_size_bytes{job="%s",mountpoint="/"}*100`, job, job, job)
		d := `max((node_filesystem_size_bytes%7Bjob%3D~"host-node"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D-node_filesystem_free_bytes%7Bjob%3D~"host-node"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D)%20*100%2F(node_filesystem_avail_bytes%20%7Bjob%3D~"host-node"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D%2B(node_filesystem_size_bytes%7Bjob%3D~"host-node"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D-node_filesystem_free_bytes%7Bjob%3D~"host-node"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D)))by(instance)`
		d6 := `max((node_filesystem_size_bytes%7Bjob%3D~"Centos6-host"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D-node_filesystem_free_bytes%7Bjob%3D~"Centos6-host"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D)%20*100%2F(node_filesystem_avail_bytes%20%7Bjob%3D~"Centos6-host"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D%2B(node_filesystem_size_bytes%7Bjob%3D~"Centos6-host"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D-node_filesystem_free_bytes%7Bjob%3D~"Centos6-host"%2Cfstype%3D~"ext.%3F%7Cxfs"%7D)))by(instance)`

		if job == "host-node" {
			disk = resource.GetDiskData(d)

		} else {
			disk = resource.GetDiskData(d6)
		}
		es := fmt.Sprintf(`node_netstat_Tcp_CurrEstab{job="%s"}`, job)
		netEs := resource.GetNetstat(es)
		tw := fmt.Sprintf(`node_sockstat_TCP_tw{job="%s"}-0`, job)
		netTW := resource.GetTCPTW(tw)
		dw := fmt.Sprintf(`max(rate(node_disk_written_bytes_total{job=~"%s"}[30s]))by(instance)/1024/1024`, job)
		DwRite := resource.GetDData(dw)
		dr := fmt.Sprintf(`max(rate(node_disk_read_bytes_total{job=~"%s"}[30s]))by(instance)/1024/1024`, job)
		Dread := resource.GetreadData(dr)
		nu := fmt.Sprintf(`max(rate(node_network_transmit_bytes_total{job="%s"}[30s])*8)by(instance)/1000/1000`, job)
		NetU := resource.GetupData(nu)
		nd := fmt.Sprintf(`max(rate(node_network_receive_bytes_total{job="%s"}[30s])*8)by(instance)/1000/1000`, job)
		NetD := resource.GetDownData(nd)
		if len(up) == len(cpu) && len(up) == len(DwRite) && len(up) == len(Dread) && len(up) == len(NetU) && len(up) == len(NetD) && len(up) == len(disk) && len(up) == len(netEs) && len(up) == len(mem) && len(up) == len(netTW) {
			CreateXlS(up, cpu, mem, disk, netEs, netTW, DwRite, Dread, NetU, NetD, job)

		}

	}

}
func main() {

	PasteData(model.C.Job.JobName...)
	fmt.Println("start is OK")
	utils.NewCrond(model.C.CronTime.CronStart, mailX.SendX)

}
