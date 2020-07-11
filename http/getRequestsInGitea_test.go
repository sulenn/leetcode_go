package http

import (
	"encoding/json"
	"github.com/sulenn/leetcode_go/modules/structs"
	"testing"
)

//测试通过项目拥有人和项目名获得仓库所有信息
func Test_GetRepositoryByOwnerAndRepo(t *testing.T) {
	body := GetData("https://gitea.com/api/v1/repos/sulenn/nudt1")
	var repository structs.Repository
	err := json.Unmarshal([]byte(body), &repository)
	if err != nil {
		t.Errorf("json 转结构体出错")
	}
	SmartPrint(repository)
}

func Test_GetCommitByOwnerRepoAndSha(t *testing.T) {
	body := GetData("https://gitea.com/api/v1/repos/sulenn/nudt1/git/commits/0ba91b33d7")
	var repo_commit structs.RepoCommit
	err := json.Unmarshal([]byte(body), &repo_commit)
	if err != nil {
		t.Errorf("json 转结构体出错")
	}
	SmartPrint(repo_commit)
}
