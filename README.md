##需求
随着算法竞赛的快速发展，各类OJ(Online Judge)系统层出不穷，题目质量也是参差不齐。而用户也存在着有的题需要去A平台上做，有的需要去B平台上做的困扰。因此，我们决定设计一个虚拟判题系统，通过爬虫技术来获取各大OJ的题目，同时实现对应的OJ的提交代理，抓取判题结果，对OJ的功能实现一个横向扩展。

---
##已完成的功能
> * oj_getter:获取OJ题目
> * ip_pool:获取代理IP
> * vjudge:代理提交,获取评测结果

---
##TODO
- [ ] 优化ip_pool,加速获取的代理IP的可用性和抓取的速度
- [ ] 支持更多的OJ


