package handlers

import (
	"encoding/json"
	"fmt"
	"mysql-app/config"
	"mysql-app/models"
	"net/http"
)

// GetUsers 获取所有用户的HTTP处理函数
func GetUsers(w http.ResponseWriter, r *http.Request) {
	db := config.GetDB()

	users, err := models.GetAllUsers(db)
	if err != nil {
		http.Error(w, fmt.Sprintf("获取用户列表失败: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// CreateUser 创建用户的HTTP处理函数
func CreateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "只支持POST请求", http.StatusMethodNotAllowed)
		return
	}

	db := config.GetDB()

	// 解析表单数据
	name := r.FormValue("name")
	email := r.FormValue("email")

	if name == "" || email == "" {
		http.Error(w, "姓名和邮箱不能为空", http.StatusBadRequest)
		return
	}

	id, err := models.CreateUser(db, name, email)
	if err != nil {
		http.Error(w, fmt.Sprintf("创建用户失败: %v", err), http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"id":      id,
		"message": "用户创建成功",
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
