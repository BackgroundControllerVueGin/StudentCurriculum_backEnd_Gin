# 接口文档

## login登录接口

使用JWT进行身份验证。如果账号密码正确，生成新的token返回给用户，如下#1。（超星学习通javaweb 4.9后端参考）

/login

```json
Post
{
    "userJobId": "3001",
    "userPassword": "123123"
}
```

```json
#1正常登录
{
	"code":200,
	"msg":"登录成功！",
	"avatar":"...（存在数据库中用户头像文件名后缀）"
	"authorities":"ROLE_student,"
	"token":"..."
}
#2密码错误或账户不存在
{
    "msg":"登陆失败",
    "code":502
}
```



## register注册接口

/register

```json
Post
{
   "userJobId":"3001",  //注册的账号必须是数据库里存在的账号 需要根据student,teacher,dean(管理员),counselor(辅导员)这四张表进行索引。可根据用户注册权限检索用户的表，防止用户越权注册. eg:注册权限为2,索引role表知道2为teacher权限，检索teacher中是否有该注册账号的信息，如果有并且尚未被注册到user表里，即可以让用户注册。
   "userPassword":"123123",
   "userRole":"3",
}
```

```json
#1正常注册
{
    "code":200,
	"msg":"注册成功",
	"data":null  //这里无所谓
}
#2用户已存在
{
	"code":201,
	"msg":"该用户已注册",
	"data":null  //这里无所谓
}
#3非法注册（权限与身份不匹配）
{
    "code":202,
	"msg":"权限不足，无法注册",
	"data":null  //这里无所谓
}
```

## public 公共资源接口

public/getUsername 主页获取并显示用户名字

```json
Post
{
    "userJobId":"3001"
}
```

```json
{
    "code": 200,
    "msg": "用户信息获取成功",
    "data": "王榄"
}
```

/public/upload 更新用户头像 (参考超星学习通javaweb 4.9后端)

```json
Post
{
    "userJobId":"3001",
    "fileList":"文件名"
}
```

/public/getMajor 根据年级 获取当年所有专业列表

```json
Post
{
    "classYear":"2022"
}
```

```json
{
    "code":"200",
    "msg":"获取所有专业",
    "data":{
        "majorList":[
            "大数据",
            "计算机科学与技术",
            "软件工程"
        ]
    }
}
```

/public/getClassName 根据年级和专业 获取到该专业所有班级

```json
Post
{
	"classYear":"2022",
	"classMajor":"软件工程"
}
```

```json
 {
     "code": 200,
     "msg": "获取到班级名列表",
     "data": {
         "classNameList": [
             "软件工程1班",
             "软件工程2班"
         ]
     }
 }
```

/public/getCourseRecordOfWeek 获取学生查询当周的课表

```json
Post
{
    "classYear": "2022",
    "classMajor": "软件工程",
    "className": "软件工程1班",
    "courseTableYear": "2023",//需要查询课表的年份
    "numberOfWeek":"1" 
}
```

```json
{
    "code": 200,
    "msg": "查找到该班级课程记录",
    "data": {
        "classYear": "2022",
        "classMajor": "软件工程",
        "className": "软件工程1班",
        "courseTableYear": "2023",
        "numberOfWeek": "1",
        "courseRecordList": [
            {
                "courserecordId": "1",
                "courserecordTableId": 202201,
                "courserecordNumberOfWeek": 1,
                "courserecordDayOfWeek": 1,
                "courserecordStartTime": "08:00:00",
                "courserecordCourseId": 1,
                "courserecordAddress": "文科类201",
                "courseList": {
                    "courseName": "软件工程",
                    "courseDeanJobId": null,
                    "courseStartDate": null,
                    "courseEndDate": null,
                    "courseTeacherJobId": "2001",
                    "courseClassId": null
                }
            },
            {
                "courserecordId": "2",
                "courserecordTableId": 202201,
                "courserecordNumberOfWeek": 1,
                "courserecordDayOfWeek": 2,
                "courserecordStartTime": "14:00:00",
                "courserecordCourseId": 2,
                "courserecordAddress": "文科类202",
                "courseList": {
                    "courseName": "软件测试",
                    "courseDeanJobId": null,
                    "courseStartDate": null,
                    "courseEndDate": null,
                    "courseTeacherJobId": "2002",
                    "courseClassId": null
                }
            },
            {
                "courserecordId": "3",
                "courserecordTableId": 202201,
                "courserecordNumberOfWeek": 1,
                "courserecordDayOfWeek": 3,
                "courserecordStartTime": "10:00:00",
                "courserecordCourseId": 3,
                "courserecordAddress": "文科类201",
                "courseList": {
                    "courseName": "数据结构",
                    "courseDeanJobId": null,
                    "courseStartDate": null,
                    "courseEndDate": null,
                    "courseTeacherJobId": "2003",
                    "courseClassId": null
                }
            },
            {
                "courserecordId": "4",
                "courserecordTableId": 202201,
                "courserecordNumberOfWeek": 1,
                "courserecordDayOfWeek": 4,
                "courserecordStartTime": "08:00:00",
                "courserecordCourseId": 4,
                "courserecordAddress": "文科类201",
                "courseList": {
                    "courseName": "人工智能",
                    "courseDeanJobId": null,
                    "courseStartDate": null,
                    "courseEndDate": null,
                    "courseTeacherJobId": "2004",
                    "courseClassId": null
                }
            },
            {
                "courserecordId": "5",
                "courserecordTableId": 202201,
                "courserecordNumberOfWeek": 1,
                "courserecordDayOfWeek": 5,
                "courserecordStartTime": "08:00:00",
                "courserecordCourseId": 5,
                "courserecordAddress": "文科类202",
                "courseList": {
                    "courseName": "c语言",
                    "courseDeanJobId": null,
                    "courseStartDate": null,
                    "courseEndDate": null,
                    "courseTeacherJobId": "2005",
                    "courseClassId": null
                }
            }
        ]
    }
}
```

## Teacher教师接口

/teacher/removeStudent

```json
Post
{
    studentJobId: "3999"
}
```

```json
#1、成功删除
{
    "code": 200,
    "msg": "学生删除成功",
    "data": null
}
#2、删除失败
{
    "code": 400,
    "msg": "该学生不存在",
    "data": null
}
```

/teacher/addStudent

```json
Post
{
	"studentJobId":"3999",
	"studentName":"ceshi",
    "studentSex":"男",
    "studentEnterYear":"2023-06-01",
    "studentMajor":"软件工程",
    "studentClassId":"1"
}
```

```json
#1、添加成功
{
    "code": 200,
    "msg": "学生添加成功",
    "data": null
}
#2、重复添加
{
    "code": 201,
    "msg": "该学生已存在",
    "data": null
}
#3、添加失败
{
    "code":400,
    "msg":"添加失败"，
    "data":null
}

```

/teacher/getStudentList

```json
Post
#success
{
	"currentPageNum":1,
	"pageSizeNum":5,
}
#fail
{
    "currentPageNum":7,
	"pageSizeNum":5,
}
```

```json
#1、成功获取
{
    "code": 200,
    "msg": "获取学生列表",
    "data": {
        "total": 22,
        "studentList": [
            {
                "studentJobId": "3001",
                "studentName": "王榄",
                "studentSex": "男",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": {
                    "classId": "1",
                    "className": "软件工程1班",
                    "classYear": "2022",
                    "classMajor": "软件工程",
                    "classCounselorJobId": null
                }
            },
            {
                "studentJobId": "3002",
                "studentName": "蔡力保",
                "studentSex": "女",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": {
                    "classId": "1",
                    "className": "软件工程1班",
                    "classYear": "2022",
                    "classMajor": "软件工程",
                    "classCounselorJobId": null
                }
            },
            {
                "studentJobId": "3004",
                "studentName": "张斯",
                "studentSex": "男",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": {
                    "classId": "1",
                    "className": "软件工程1班",
                    "classYear": "2022",
                    "classMajor": "软件工程",
                    "classCounselorJobId": null
                }
            },
            {
                "studentJobId": "3008",
                "studentName": "周杰",
                "studentSex": "男",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": {
                    "classId": "1",
                    "className": "软件工程1班",
                    "classYear": "2022",
                    "classMajor": "软件工程",
                    "classCounselorJobId": null
                }
            },
            {
                "studentJobId": "3009",
                "studentName": "刘一条",
                "studentSex": "男",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": {
                    "classId": "1",
                    "className": "软件工程1班",
                    "classYear": "2022",
                    "classMajor": "软件工程",
                    "classCounselorJobId": null
                }
            }
        ]
    }
}
#2、获取失败
{
    "code": 201,
    "msg": "获取学生列表失败",
    "data": null
}
```

/teacher/getStudent

```json
Post
#success
{
	"studentJobId":"3001"
}
#fail
{
  "studentJobId":"3099"  
}
```

```json
#1、成功获取学生信息
{
    "code": 200,
    "msg": "获取学生信息成功",
    "data": {
        "studentList": [
            {
                "studentJobId": "3001",
                "studentName": "王榄",
                "studentSex": "男",
                "studentEnterYear": "2023-05-26",
                "studentMajor": "软件工程",
                "studentClassId": "1",
                "studentClass": null
            }
        ]
    }
}
#2、获取学生信息失败
{
    "code": 201,
    "msg": "获取学生信息失败",
    "data": null
}
```



## Student学生接口

/student/getConselor

```json
Post
{
   "studentJobId":"3001"
}
```

```json
#1、成功获取学生辅导员信息
{
    "code": 200,
    "msg": "查询成功",
    "data": [
        {
            "counselorJobId": "4001",
            "counselorName": "辅导员一号",
            "counselorSex": "男"
        }
    ]
}
```

/student/submitApplication

```json
Post(换成学生的token)
{
    "applicationApplicantJobId": "3001",
    "applicationContent": "ceshiceshi",
    "applicationStartTime": "2023-06-01 00:00:00",
    "applicationEndTime": "2023-06-08 00:00:00",
    "applicationReviewerJobId": "4001",
    "applicationTypeId": "1" 
}
```

```json
#1、提交成功
{
    "code": 200,
    "msg": "提交申请成功",
    "data": null
}
```

## counselor辅导员接口

/counselor/updateApplication

```json
Post(换成辅导员的token)
{
	"applicationId": 20,
	"applicationReply": "",
	"applicationStatus": "rejected"
}
```

```json
#1、审核完成
{
    "code": 200,
    "msg": "审核完成",
    "data": null
}
```

/counselor/getApplicationList

```json
Post(换成辅导员的token)
{
	userJobId: "4001"
}
```

```json
#1、成功获取
{
    "code": 200,
    "msg": "获取该老师的申请表",
    "data": [
        {
            "applicationId": 3,
            "applicationTypeId": 1,
            "applicationContent": "ceshi",
            "applicationCreatedTime": "2023-06-05T02:08:13",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-04T12:57:35",
            "applicationEndTime": "2023-06-05T12:57:39",
            "applicationStatus": "approved",
            "applicationReply": "拒绝asdf",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "病假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 14,
            "applicationTypeId": 1,
            "applicationContent": "外出游玩1111",
            "applicationCreatedTime": "2023-06-12T15:17:50",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-01T00:00:00",
            "applicationEndTime": "2023-06-20T00:00:00",
            "applicationStatus": "rejected",
            "applicationReply": "不同意 时间太长",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "病假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 20,
            "applicationTypeId": 1,
            "applicationContent": "ceshiceshi",
            "applicationCreatedTime": "2023-06-27T18:17:18",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-01T00:00:00",
            "applicationEndTime": "2023-06-08T00:00:00",
            "applicationStatus": "rejected",
            "applicationReply": "",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "病假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 21,
            "applicationTypeId": 1,
            "applicationContent": "ceshiceshi",
            "applicationCreatedTime": "2023-06-27T18:20:35",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-01T00:00:00",
            "applicationEndTime": "2023-06-08T00:00:00",
            "applicationStatus": "rejected",
            "applicationReply": "ceshi",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "病假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 13,
            "applicationTypeId": 2,
            "applicationContent": "ceshiceshi",
            "applicationCreatedTime": "2023-06-12T00:18:34",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-12T00:00:00",
            "applicationEndTime": "2023-06-14T00:00:00",
            "applicationStatus": "rejected",
            "applicationReply": "",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 15,
            "applicationTypeId": 2,
            "applicationContent": "生病生病生病",
            "applicationCreatedTime": "2023-06-25T23:13:34",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-26T00:00:00",
            "applicationEndTime": "2023-06-27T00:00:00",
            "applicationStatus": "approved",
            "applicationReply": "特殊的纺纱机快点发货就开始",
            "applicationApplicantJobId": "3013",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "李红微",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 16,
            "applicationTypeId": 2,
            "applicationContent": "士大夫艰苦拉到JFK双方都是",
            "applicationCreatedTime": "2023-06-25T23:26:40",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-26T00:00:00",
            "applicationEndTime": "2023-06-27T00:00:00",
            "applicationStatus": "approved",
            "applicationReply": "塑料袋放进撒旦开发",
            "applicationApplicantJobId": "3010",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "驲期",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 17,
            "applicationTypeId": 2,
            "applicationContent": "是浪费大家撒扩大飞机",
            "applicationCreatedTime": "2023-06-25T23:33:17",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-26T00:00:00",
            "applicationEndTime": "2023-06-27T00:00:00",
            "applicationStatus": "approved",
            "applicationReply": "1111111",
            "applicationApplicantJobId": "3013",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "李红微",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 18,
            "applicationTypeId": 2,
            "applicationContent": "6161616",
            "applicationCreatedTime": "2023-06-26T01:22:37",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-01T00:00:00",
            "applicationEndTime": "2023-06-03T00:00:00",
            "applicationStatus": "approved",
            "applicationReply": "注意安全玩的开心",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        },
        {
            "applicationId": 19,
            "applicationTypeId": 2,
            "applicationContent": "1231231",
            "applicationCreatedTime": "2023-06-27T05:12:40",
            "applicationReviewedTime": null,
            "applicationStartTime": "2023-06-06T00:00:00",
            "applicationEndTime": "2023-06-14T00:00:00",
            "applicationStatus": "rejected",
            "applicationReply": "1111111",
            "applicationApplicantJobId": "3001",
            "applicationReviewerJobId": "4001",
            "applicationType": {
                "applicationTypeId": 0,
                "applicationTypeName": "事假"
            },
            "student": {
                "studentJobId": null,
                "studentName": "王榄",
                "studentSex": null,
                "studentEnterYear": null,
                "studentMajor": null,
                "studentClassId": null,
                "studentClass": null
            }
        }
    ]
}
```

