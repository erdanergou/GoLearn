package main

/*Viper
Viper是适用于Go应用程序（包括Twelve-Factor App）的完整配置解决方案。它被设计用于在应用程序中工作，并且可以处理所有类型的配置需求和格式

Viper会按照下面的优先级。每个项目的优先级都高于它下面的项目:

	显示调用Set设置值
	命令行参数（flag）
	环境变量
	配置文件
	key/value存储
	默认值
重要： 目前Viper配置的键（Key）是大小写不敏感的。目前正在讨论是否将这一选项设为可选

读取配置文件

Viper需要最少知道在哪里查找配置文件的配置
Viper可以搜索多个路径，但目前单个Viper实例只支持单个配置文件。Viper不默认任何配置搜索路径，将默认决策留给应用程序
*/

func main() {
	viper.
}
