package vjudge

import "testing"

func TestCodeVSJudger_Submit(t *testing.T) {
	c := CodeVSJudger{}
	c.Submit("1000", "C++", "#include<iostream>using namespace std;int main(){int a,b;while(cin>>a>>b){cout<<a+b<<endl;}return 0;}")
}
