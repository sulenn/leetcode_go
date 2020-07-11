package http

import (
	"testing"
)

//测试通过项目拥有人和项目名获得仓库所有信息
func Test_GetCompareTwoCommits(t *testing.T) {
	GetData("https://api.github.com/repos/sulenn/blogDir/compare/9ed414d812f0c969b179b594a352932b01210c50...f1b2c8520cef15da81c09f9faddc384b633cda07")
	//GetData("https://api.github.com/repos/sulenn/blogDir/commits/f1b2c8520cef15da81c09f9faddc384b633cda07")
}
