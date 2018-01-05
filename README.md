# 数据分析与查重系统

##功能
1.对代码进行判重(英文)<br/>
2.使用数据挖掘相关知识进行关联匹配推荐<br/>
3.对相关数据进行统计<br/>
4.对异常数据预警<br/>

---
##难点
1.高并发,实时性<br/>
并发主要是在比赛时同时处理多用户代码分析,同时需要实时显示,对异常数据需要及时反馈<br/>
2.关联算法<br/>
使用关联算法实现相似数据的推荐<br/>

---
##使用技术
使用Go处理高并发,redis缓存实现实时性

