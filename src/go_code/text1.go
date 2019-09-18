package main
// import (
// 	"fmt"
// 	"os"
// 	"bufio"
// 	"io"
// )
/*
	文件的基本介绍:文件是指针类型
		- 流: 数据在数据源(文件)和程序(内存)之间经历的路径
		- 输入流: 数据从数据源(文件)到程序(内存)的路径  ---->读文件
		- 输出流: 数据从程序(内存)到数据源(文件)的路径  ---->写文件

	应用的文件os.File, File是结构体,里面有读写方法,File代表一个文件句柄
	打开文件和关闭文件
		os.File.Open()  打开
		os.File.Close()  关闭
	例子
		func main() {

			file, err := os.Open("/home/beriuta/Desktop/weekly_plan.txt")
			if err != nil {
				fmt.Println("打开文件错误:", err)
			} 
			// 输出文件是什么
			fmt.Println(file)  // &{0xc000092120}
			// 关闭文件
			err = file.Close()
			if err != nil {
				fmt.Println("关闭是错误:", err)
			}

		}

	读取文件操作应用实例
		读取文件内容并显示在终端上(带缓冲区的方式),使用os.Open(), file.Close(), bufio.NewReader(), reader.Readstring函数和方法
	
	示例:
		func main() {
			file, err := os.Open("/home/beriuta/Desktop/weekly_plan.txt")
			if err != nil {
				fmt.Println("打开文件错误:", err)
			} 
			
			// 当函数退出时,要及时的关闭file.用defer,最后才执行
			defer file.Close()

			// 创建一个  *Reader 是带缓冲区的
			// 在官方文档中,说明这个缓冲区为4096
			// const {
				// defaultBufSize = 4096
			// }
			reader := bufio.NewReader(file)
			for {
				str, err := reader.ReadString('\n')  // 括号里面表示读到什么地方结束
				if err == io.EOF {   // io.EOF表示文件的末尾
					break
				}
				// 输出内容
				fmt.Print(str)
			}
			fmt.Println("读取文件结束.........")
			
		}

	读取文件操作应用实例
		读取文件的内容并显示在终端(使用ioutil一次将整个文件读入到内存中),这种方式适用与于
		文件不打的情况,相关方法和函数(ioutil.ReadFile)
	示例:
		func main() {
			file := "/home/beriuta/Desktop/weekly_plan.txt"
			content, err := ioutil.ReadFile(file)  // 返回的是一个[]byte, 这里不需要打开或者关闭,因为这两个方法都被封装到这个方法了
			if err != nil {
				fmt.Printf("报错信息是: %v\n", err)
			}

			// 那读取到的内容显示到终端
			fmt.Printf("%v\n", content)
			// [55 46 50 57 32 229 145 168 232 174 161 229 136 146 10 49 46 32 230 155 180 230 150 176 231 148 168 232 189 166 231 187 159 232 174 161 232 161 168 228 187 165 229 143 138 230 138 165 232 173 166 228 191 161 230 129 175 231 154 132 229 137 141 231 171 175 230 150 135 230 161 163 40 229 183 178 229 174 140 230 136 144 41 10 50 46 32 231 188 150 229 134 153 230 138 165 232 173 166 228 191 161 230 129 175 231 154 132 233 128 187 232 190 145 40 232 182 133 233 128 159 41 40 229 183 178 229 174 140 230 136 144 41 10 51 46 32 228 184 137 228 184 170 232 180 166 229 143 183 44 230 160 185 230 141 174 228 184 141 229 144 140 231 154 132 232 180 166 229 143 183 232 167 146 232 137 178 230 157 165 229 177 149 231 164 186 228 184 141 229 144 140 231 154 132 229 134 133 229 174 185 40 230 173 164 228 184 137 228 184 170 232 180 166 229 143 183 228 184 186 232 175 149 231 148 168 232 180 166 229 143 183 44 229 143 170 232 131 189 230 159 165 231 156 139 228 191 161 230 129 175 44 228 184 141 232 131 189 228 191 174 230 148 185 41 32 32 40 229 183 178 229 174 140 230 136 144 41 10 52 46 32 230 160 185 230 141 174 231 148 168 232 189 166 228 186 186 231 154 132 232 186 171 228 187 189 44 230 183 187 229 138 160 229 187 186 232 174 174 230 180 190 232 189 166 230 136 150 232 128 133 228 188 152 229 133 136 230 180 190 232 189 166 231 154 132 230 143 144 233 134 146 32 32 40 229 183 178 229 174 140 230 136 144 41 10 53 46 32 230 160 185 230 141 174 79 66 68 229 141 143 232 174 174 228 184 138 228 188 160 231 154 132 233 169 190 233 169 182 228 191 161 230 129 175 40 232 182 133 233 128 159 44 230 128 165 229 138 160 233 128 159 44 230 128 165 229 135 143 233 128 159 44 230 128 165 232 189 172 229 188 175 41 44 231 187 153 233 169 190 233 169 182 229 145 152 232 175 132 228 187 183 10 10 32]
			// 如果想要转string类型
			fmt.Printf("%v\n", string(content))  // 输出的就是字符串类型
		}

	写文件操作应用实例
	基本介绍
		func OpenFile(name string, flag int, perm FileModel) (file *File, err error)
		第一个参数是文件名,第二个参数是文件打开模式,第三个参数是权限控制:r---4, w---2, x---1可执行
	基本应用实例-方式1
	1. 创建一个新文件,写入内容5句"hello,world!!"使用os.OpenFile(),bufio.NewWriterr()
		func main() {
			// 创建一个新文件,写入内容
			filePath := "/home/beriuta/Desktop/abc.txt"
			file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_CREATE, 0666)  // 表示如果没有此文件就新建一个此文件并且写入内容
			if err != nil {
				fmt.Printf("写入错误:", err)
				return
			}
			defer file.Close()
			// 准备写入5句话
			str := "helle world\r\n"  // 有些编辑器认识\r有些认识\n
			// 写入时,使用带缓存的 *Writer
			writer := bufio.NewWriter(file)
			for i:= 0; i < 5; i++ {
				writer.WriteString(str)
			}

			// 因为writer是带缓存的,因此在调用writerString方法时,其实内容是直接写到缓存中
			// 如果想要写到文件中,需要调用一个方法
			writer.Flush()  // 将数据写入文件中,如果不加这句,会造成数据丢失
		}
	2. 打开一个存在的文件中,将原来的内容覆盖成新的内容10句"你好!小岳岳"
		func main() {
			// 给出一个已有的文件,覆盖写入内容
			filePath := "/home/beriuta/Desktop/abc.txt"
			file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_TRUNC, 0666)  // 表示如果没有此文件就新建一个此文件并且写入内容
			if err != nil {
				fmt.Printf("写入错误:", err)
				return
			}
			defer file.Close()
			// 准备写入5句话
			str := "你好!小岳岳\r\n"  // 有些编辑器认识\r有些认识\n
			// 写入时,使用带缓存的 *Writer
			writer := bufio.NewWriter(file)
			for i:= 0; i < 10; i++ {
				writer.WriteString(str)
			}

			// 因为writer是带缓存的,因此在调用writerString方法时,其实内容是直接写到缓存中
			// 如果想要写到文件中,需要调用一个方法
			writer.Flush()  // 将数据写入文件中,如果不加这句,会造成数据丢失
		}
	3. 打开一个存在的文件,在原来的内容追加内容"ABCI SENGING"
		func main() {
			// 给出一个已有的文件,追加写入内容
			filePath := "/home/beriuta/Desktop/abc.txt"
			file, err := os.OpenFile(filePath, os.O_WRONLY | os.O_APPEND, 0666)  // 表示如果没有此文件就新建一个此文件并且写入内容
			if err != nil {
				fmt.Printf("写入错误:", err)
				return
			}
			defer file.Close()
			// 准备写入5句话
			str := "我是追加的内容,一共有三句\r\n"  // 有些编辑器认识\r有些认识\n
			// 写入时,使用带缓存的 *Writer
			writer := bufio.NewWriter(file)
			for i:= 0; i < 3; i++ {
				writer.WriteString(str)
			}

			// 因为writer是带缓存的,因此在调用writerString方法时,其实内容是直接写到缓存中
			// 如果想要写到文件中,需要调用一个方法
			writer.Flush()  // 将数据写入文件中,如果不加这句,会造成数据丢失
		}
	4. 打开一个存在的文件,将原来的内容读出显示在终端,并且追加5句"hello! 深圳"
		func main() {
			// 打开一个存在的文件,将原来的内容读出显示在终端,并且追加5句"hello! 深圳"
			// 既有读的功能,又有写的功能
			filePath := "/home/beriuta/Desktop/abc.txt"
			file, err := os.OpenFile(filePath, os.O_RDWR | os.O_APPEND, 0666)  // 表示如果没有此文件就新建一个此文件并且写入内容
			if err != nil {
				fmt.Printf("写入错误:", err)
				return
			}
			defer file.Close()

			// 先读出来
			reader := bufio.NewReader(file)
			for {
				str, err := reader.ReadString('\n')
				if err == io.EOF {
					break
				}
				fmt.Print(str)
			}

			// 准备写入5句话
			str := "hello! 深圳\r\n"  // 有些编辑器认识\r有些认识\n
			// 写入时,使用带缓存的 *Writer
			writer := bufio.NewWriter(file)
			for i:= 0; i < 5; i++ {
				writer.WriteString(str)
			}

			// 因为writer是带缓存的,因此在调用writerString方法时,其实内容是直接写到缓存中
			// 如果想要写到文件中,需要调用一个方法
			writer.Flush()  // 将数据写入文件中,如果不加这句,会造成数据丢失
		}

	写文件操作应用实例
	将一个文件的内容,写入另一个文件中,两个文件必须要存在才能操作
	方法 : ioutil.ReadFile  ioutil.WriteFile
		func main() {
			filePath1 := "/home/beriuta/Desktop/weekly_plan.txt"
			filePath2 := "/home/beriuta/Desktop/abc.txt"
			// 先读取1的数据,然后写入2的文件中
			content, err := ioutil.ReadFile(filePath1)
			if err != nil {
				// 说明文件有错误
				fmt.Println("读文件出错了:", err)
				return
			}
			err = ioutil.WriteFile(filePath2, content, 06666)
			if err != nil {
				fmt.Println("写文件出错了:", err)
				return
			}
		}
	
	判断文件或文件夹是否存在的方法:os.Stat()函数返回的错误值进行判断
		1. 如果返回的错误为nil说明文件或文件夹存在
		2. 如果返回的错误类型使用os.IsNotExist()判断为True,说明文件不存在
		3. 如果返回错误的为其他类型,则不确定是否存在

	拷贝文件
		将一张图片/电影/MP3拷贝到另一个文件, 使用的是io包
		func Copy(dst Writer, src Reader)(written int64, err error)
		注意: Copy函数是io包提供的
		// 自己编写一个方法, 只能拷贝除了txt文件的内容
		func fileCopy(dst string, src string) (written int64, err error){
			srcFile, err := os.Open(src)
			if err != nil {
				fmt.Println("打开文件失败:", err)
				return
			}
			defer srcFile.Close()
			// 通过src,获取到Reader
			reader := bufio.NewReader(srcFile)
			// 打开dst文件,因为这个文件要写,并且不一定有这个路径存在,所以不能直接用os.Open()
			dstFile, err := os.OpenFile(dst, os.O_WRONLY | os.O_CREATE, 0666)
			if err != nil {
				fmt.Println("写入文件错误:", err)
				return
			}
			defer dstFile.Close()
			// 通过dst,获取到Writer
			writer := bufio.NewWriter(dstFile)
			return io.Copy(writer, reader)
		}
		func main() {
			// 图片的路径跟复制图片的路径(这个可以是假的,如果程序判断不存在这个路径将会新建)
			filePath := "/home/beriuta/Desktop/abc.txt"
			writerFile := "/home/beriuta/Desktop/ffff.txt"
			_, err := fileCopy(writerFile, filePath)
			if err != nil {
				fmt.Println("拷贝错误:", err)
			} else {
				fmt.Println("拷贝完成")
			}
		}

	文件编程应用实例
	统计英文,数字,空格和其他字符数量
*/














	


