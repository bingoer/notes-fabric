0 前提 
 fabric 工程已启动
 ./network_setup.sh up

1 下载java-sdk官方代码
https://github.com/hyperledger/fabric-sdk-java/tree/release-1.2
git clone -b release-1.2 https://github.com/hyperledger/fabric-sdk-java.git
git clone -b release-1.2  https://gitee.com/xugy/fabric-sdk-java-object.git

2 修改pom.xml
在properties中加入下面一句话：
<os.detected.classifier>windows-x86_64</os.detected.classifier>

3 导入maven工程 
--编译工程
mvn install 

--重新生成
--删掉.project 和 .classpath文件 
--重新生成
mvn eclipse:eclipse

--导入工程
确保maven 命令 同eclipse 


https://gitee.com/xugy/fabric-sdk-java-object.git


4 运行工程
--修改配置
TestConfig.java
1. LOCALHOST 对应IP 

	
5 运行测试
	End2endIT.java


6 注意事项
 1) 时间同服务器时间保持一致



https://github.com/xiangxingchina/study.git


 x

Failed to read the project description file (.project) for 'fabric-sdk-java'.  The file has been changed on disk, and it now contains invalid information.  The project will not function properly until the description file is restored to a valid state.
