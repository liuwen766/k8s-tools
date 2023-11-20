- 笔记-基础篇-1(P1-P28)：[https://blog.csdn.net/hancoder/article/details/106922139](https://blog.csdn.net/hancoder/article/details/106922139)
- 笔记-基础篇-2(P28-P100)：[https://blog.csdn.net/hancoder/article/details/107612619](https://blog.csdn.net/hancoder/article/details/107612619)
- 笔记-高级篇(P340)：[https://blog.csdn.net/hancoder/article/details/107612746](https://blog.csdn.net/hancoder/article/details/107612746)
- 笔记-vue：[https://blog.csdn.net/hancoder/article/details/107007605](https://blog.csdn.net/hancoder/article/details/107007605)
- 笔记-elastic search、上架、检索：[https://blog.csdn.net/hancoder/article/details/113922398](https://blog.csdn.net/hancoder/article/details/113922398)
- 笔记-认证服务：[https://blog.csdn.net/hancoder/article/details/114242184](https://blog.csdn.net/hancoder/article/details/114242184)
- 笔记-分布式锁与缓存：[https://blog.csdn.net/hancoder/article/details/114004280](https://blog.csdn.net/hancoder/article/details/114004280)
- 笔记-集群篇：[https://blog.csdn.net/hancoder/article/details/107612802](https://blog.csdn.net/hancoder/article/details/107612802)
- springcloud笔记：[https://blog.csdn.net/hancoder/article/details/109063671](https://blog.csdn.net/hancoder/article/details/109063671)
- 笔记版本说明：2020年提供过笔记文档，但只有P1-P50的内容，2021年整理了P340的内容。请点击标题下面分栏查看系列笔记
- 声明：

  - 可以白嫖，但请勿转载发布，笔记手打不易
  - 本系列笔记不断迭代优化，csdn：hancoder上是最新版内容，10W字都是在csdn免费开放观看的。
  - 离线md笔记文件获取方式见文末。2021-3版本的md笔记打完压缩包共500k（云图床），包括本项目笔记，还有cloud、docker、mybatis-plus、rabbitMQ等个人相关笔记
- sql：[https://github.com/FermHan/gulimall/sql文件](https://github.com/FermHan/gulimall/sql文件)
- 本项目其他笔记见专栏：[https://blog.csdn.net/hancoder/category_10822407.html](https://blog.csdn.net/hancoder/category_10822407.html)

> 请直接ctrl+F搜索内容

本文为谷粒商城项目而简要学习

前后端对比

<img src="https://i0.hdslb.com/bfs/album/3a64bb02b0ece6d0fdc6b3606859246b15b0872e.png" style="zoom:33%;" />

# 一、VS环境搭建

之前写过了，需要在VScode中安装如下插件：

- Auto Close Tag
- Auto Rename Tag
- Chinese
- ESlint
- HTML CSS Support
- HTML Snippets
- JavaScript ES6
- Live Server
- open in brower
- Vetur33
- Vue 2Snippets 语法提示
- 谷歌浏览器中安装插件Vue Devtools

# 二、ES6

`ECMAScript6.0`（以下简称ES6，ECMAScript是一种由Ecma国际通过ECMA-262标准化的脚本），是JavaScript语言的下一代标准，2015年6月正式发布，从ES6开始的版本号采用年号，如

- ES2015，就是ES6。
- ES2016，就是ES7。
- ES2017，就是ES8。

> ECMAScript是规范，JS的规范的具体实现。

打开VSCode---打开文件夹---新建es6文件夹---新建文件1、let.html---`shift+!+Enter`生成模板。填入下面内容后，右键open with live server

##### 1、let作用域.html

- let不会作用到{}外，var会越域跳到{}外
- var可以多次声明同一变量，let会报错
- var定义之前可以使用，let定义之前不可使用。（变量提升问题）

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    

    <script>
       // var 声明的变量往往会越域
       // let 声明的变量有严格局部作用域
         {
             var a = 1;
             let b = 2;
         }
         console.log(a);  // 1
         console.log(b);  // ReferenceError: b is not defined

         // var 可以声明多次
         // let 只能声明一次
         var m = 1
         var m = 2
         let n = 3
//         let n = 4
         console.log(m)  // 2
         console.log(n)  // Identifier 'n' has already been declared

        // var 会变量提升
        // let 不存在变量提升
         console.log(x);  // undefined
         var x = 10;
         console.log(y);   //ReferenceError: y is not defined
         let y = 20;

        // let
        // 1. const声明之后不允许改变
           // 2. 一但声明必须初始化，否则会报错
        const a = 1;
        a = 3; //Uncaught TypeError: Assignment to constant variable.
    
    </script>

</body>
</html>
```

- var在{}之外也起作用
- let在{}不起作用
- var多次声明同一变量不会报错，let多次声明会报错，只能声明一次。
- var 会变量提升（打印和定义可以顺序反）。let 不存在变量提升（顺序不能反）
- let的·声明之后不允许改变

##### 2、解构表达式.html

- 支持`let arr = [1,2,3];     let [a,b,c] = arr;`这种语法
- 支持对象解析：`const { name: abc, age, language } = person;`    冒号代表改名，`旧:新`
- 字符串函数
- 支持一个字符串为多行``
- 占位符功能 ${}

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>

    <script>
        //数组解构
        let arr = [1,2,3];
        // // let a = arr[0];
        // // let b = arr[1];
        // // let c = arr[2];

        let [a,b,c] = arr;
        console.log(a,b,c)

        const person = {
            name: "jack",
            age: 21,
            language: ['java', 'js', 'css']
        }
        //         const name = person.name;
        //         const age = person.age;
        //         const language = person.language;

        //对象解构 // 把name属性变为abc，声明了abc、age、language三个变量
        const { name: abc, age, language } = person;
        console.log(abc, age, language)

        //4、字符串扩展
        let str = "hello.vue";
        console.log(str.startsWith("hello"));//true
        console.log(str.endsWith(".vue"));//true
        console.log(str.includes("e"));//true
        console.log(str.includes("hello"));//true

        //字符串模板 ``可以定义多行字符串
        let ss = `<div>
                    <span>hello world<span>
                </div>`;
        console.log(ss);
        
        function fun() {
            return "这是一个函数"
        }

        // 2、字符串插入变量和表达式。变量名写在 ${} 中，${} 中可以放入 JavaScript 表达式。
        let info = `我是${abc}，今年${age + 10}了, 我想说： ${fun()}`;
        console.log(info);

    </script>
</body>
</html>
```



##### 3、函数优化.html

- 支持函数形参默认值  `function add2(a, b = 1) {`
- 支持不定参数  `function fun(...values) {`
- 支持箭头函数 `var print = obj => console.log(obj);`
- 箭头函数+结构 `var hello2 = ({name}) => console.log("hello," +name);`，本来应该是person.name

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>

    <script>
        //在ES6以前，我们无法给一个函数参数设置默认值，只能采用变通写法：
        function add(a, b) {
            // 判断b是否为空，为空就给默认值1
            b = b || 1;
            return a + b;
        }
        // 传一个参数
        console.log(add(10));


        //现在可以这么写：直接给参数写上默认值，没传就会自动使用默认值
        function add2(a, b = 1) {
            return a + b;
        }
        console.log(add2(20));


        //2）、不定参数
        function fun(...values) {
            console.log(values.length)
        }
        fun(1, 2)      //2
        fun(1, 2, 3, 4)  //4

        //3）、箭头函数。lambda
        //以前声明一个方法
        // var print = function (obj) {
        //     console.log(obj);
        // }
        var print = obj => console.log(obj);
        print("hello");

        var sum = function (a, b) {
            c = a + b;
            return a + c;
        }

        var sum2 = (a, b) => a + b;
        console.log(sum2(11, 12));

        var sum3 = (a, b) => {
            c = a + b;
            return a + c;
        }
        console.log(sum3(10, 20))


        const person = {
            name: "jack",
            age: 21,
            language: ['java', 'js', 'css']
        }

        function hello(person) {
            console.log("hello," + person.name)
        }

        //箭头函数+解构
        var hello2 = ({name}) => console.log("hello," +name);
        hello2(person);

    </script>
</body>
</html>
```

​	

##### 4、对象优化.html

对象简写很常见

- 可以获取map的键值对等`Object.keys()`、values、entries
- Object.assgn(target,source1,source2)   合并
- 如果属性名和属性值的变量名相同，可以省略`const person2 = { age, name }` 代表age属性的值是变量age的值
- ...代表取出该对象所有属性拷贝到当前对象。let someone = { ...p1 }


```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    <script>
        const person = {
            name: "jack",
            age: 21,
            language: ['java', 'js', 'css']
        }

        console.log(Object.keys(person));//["name", "age", "language"]
        console.log(Object.values(person));//["jack", 21, Array(3)]
        console.log(Object.entries(person));//[Array(2), Array(2), Array(2)]

        const target  = { a: 1 };
        const source1 = { b: 2 };
        const source2 = { c: 3 };

        // 合并
        //{a:1,b:2,c:3}
        Object.assign(target, source1, source2);

        console.log(target);//["name", "age", "language"]

        //2）、声明对象简写
        const age = 23
        const name = "张三"
        const person1 = { age: age, name: name }
        // 等价于
        const person2 = { age, name }//声明对象简写
        console.log(person2);

        //3）、对象的函数属性简写
        let person3 = {
            name: "jack",
            // 以前：
            eat: function (food) {
                console.log(this.name + "在吃" + food);
            },
            //箭头函数this不能使用，要使用的话需要使用：对象.属性
            eat2: food => console.log(person3.name + "在吃" + food),
            eat3(food) {
                console.log(this.name + "在吃" + food);
            }
        }

        person3.eat("香蕉");
        person3.eat2("苹果")
        person3.eat3("橘子");

        //4）、对象拓展运算符

        // 1、拷贝对象（深拷贝）
        let p1 = { name: "Amy", age: 15 }
        let someone = { ...p1 }
        console.log(someone)  //{name: "Amy", age: 15}

        // 2、合并对象
        let age1 = { age: 15 }
        let name1 = { name: "Amy" }
        let p2 = {name:"zhangsan"}
        p2 = { ...age1, ...name1 } 
        console.log(p2)
    </script>
</body>

</html>
```



##### 5、map和reduce.html
`map()`：接收一个函数，将原数组中的所有元素用这个函数处理后放入新数组返回。
`reduce()` 为数组中的每一个元素依次执行回调函数，不包括数组中被删除或从未被赋值的元素，

```JS
arr = arr.map(item=> item*2);

//[2, 40, -10, 6]
arr.reduce(callback,[initialValue])
/**
    1、previousValue （上一次调用回调返回的值，或者是提供的初始值（initialValue））
    2、currentValue （数组中当前被处理的元素）
    3、index （当前元素在数组中的索引）
    4、array （调用 reduce 的数组）*/
```

- map处理，arr = arr.map(item=> item*2);
- reduce。arr.reduce((原来的值,处理后的值即return的值)=>{

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
    
    <script>
        //数组中新增了map和reduce方法。
         let arr = ['1', '20', '-5', '3'];
         
        //map()：接收一个函数，将原数组中的所有元素用这个函数处理后放入新数组返回。
        //  arr = arr.map((item)=>{
        //     return item*2
        //  });
         arr = arr.map(item=> item*2);

        

         console.log(arr);
        //reduce() 为数组中的每一个元素依次执行回调函数，不包括数组中被删除或从未被赋值的元素，
        //[2, 40, -10, 6]
        //arr.reduce(callback,[initialValue])
        /**
        1、previousValue （上一次调用回调返回的值，或者是提供的初始值（initialValue））
        2、currentValue （数组中当前被处理的元素）
        3、index （当前元素在数组中的索引）
        4、array （调用 reduce 的数组）*/
        let result = arr.reduce((a,b)=>{
            console.log("上一次处理后："+a);
            console.log("当前正在处理："+b);
            return a + b;
        },100);
        console.log(result)

    
    </script>
</body>
</html>
```

##### 6、==promise.html==

优化异步操作。封装ajax

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
    <script src="https://cdn.bootcss.com/jquery/3.4.1/jquery.min.js"></script>
</head>
<body>
    <script>
        //1、查出当前用户信息
        //2、按照当前用户的id查出他的课程
        //3、按照当前课程id查出分数
        // $.ajax({
        //     url: "mock/user.json",
        //     success(data) {
        //         console.log("查询用户：", data);
        //         $.ajax({
        //             url: `mock/user_corse_${data.id}.json`,
        //             success(data) {
        //                 console.log("查询到课程：", data);
        //                 $.ajax({
        //                     url: `mock/corse_score_${data.id}.json`,
        //                     success(data) {
        //                         console.log("查询到分数：", data);
        //                     },
        //                     error(error) {
        //                         console.log("出现异常了：" + error);
        //                     }
        //                 });
        //             },
        //             error(error) {
        //                 console.log("出现异常了：" + error);
        //             }
        //         });
        //     },
        //     error(error) {
        //         console.log("出现异常了：" + error);
        //     }
        // });


        //1、Promise可以封装异步操作
        // let p = new Promise((resolve, reject) => { //传入成功解析，失败拒绝
        //     //1、异步操作
        //     $.ajax({
        //         url: "mock/user.json",
        //         success: function (data) {
        //             console.log("查询用户成功:", data)
        //             resolve(data);
        //         },
        //         error: function (err) {
        //             reject(err);
        //         }
        //     });
        // });

        // p.then((obj) => { //成功以后做什么
        //     return new Promise((resolve, reject) => {
        //         $.ajax({
        //             url: `mock/user_corse_${obj.id}.json`,
        //             success: function (data) {
        //                 console.log("查询用户课程成功:", data)
        //                 resolve(data);
        //             },
        //             error: function (err) {
        //                 reject(err)
        //             }
        //         });
        //     })
        // }).then((data) => { //成功以后干什么
        //     console.log("上一步的结果", data)
        //     $.ajax({
        //         url: `mock/corse_score_${data.id}.json`,
        //         success: function (data) {
        //             console.log("查询课程得分成功:", data)
        //         },
        //         error: function (err) {
        //         }
        //     });
        // })

        function get(url, data) { //自己定义一个方法整合一下
            return new Promise((resolve, reject) => {
                $.ajax({
                    url: url,
                    data: data,
                    success: function (data) {
                        resolve(data);
                    },
                    error: function (err) {
                        reject(err)
                    }
                })
            });
        }

        get("mock/user.json")
            .then((data) => {
                console.log("用户查询成功~~~:", data)
            	//  new Promise
                return get(`mock/user_corse_${data.id}.json`);
            })
            .then((data) => {
                console.log("课程查询成功~~~:", data)
                return get(`mock/corse_score_${data.id}.json`);
            })
            .then((data)=>{
                console.log("课程成绩查询成功~~~:", data)
            })
            .catch((err)=>{ //失败的话catch
                console.log("出现异常",err)
            });

    </script>
</body>

</html>
```

corse_score_10.json 得分

```json
{
    "id": 100,
    "score": 90
}
```

user.json 用户

```json
{
    "id": 1,
    "name": "zhangsan",
    "password": "123456"
}
```

user_corse_1.json 课程

```json
{
    "id": 10,
    "name": "chinese"
}
```

以前嵌套ajax的时候很繁琐。

- 把Ajax封装到Promise中，赋值给let p
- 在Ajax中成功使用resolve(data)，失败使用reject(err)
- p.then().catch()

### ==模块化import/export==

模块化就是把代码进行拆分，方便重复利用。类似于java中的导包，而JS换了个概念，是导模块。

模块功能主要有两个命令构成 export 和import

- export用于规定模块的对外接口
- import用于导入其他模块提供的功能

user.js

```js
var name = "jack"
var age = 21
function add(a,b){
    return a + b;
}
// 导出变量和函数
export {name,age,add}
```

hello.js

```js
// export const util = {
//     sum(a, b) {
//         return a + b;
//     }
// }

// 导出后可以重命名
export default {
    sum(a, b) {
        return a + b;
    }
}
// export {util}

//`export`不仅可以导出对象，一切JS变量都可以导出。比如：基本类型变量、函数、数组、对象。
```

main.js

```js
import abc from "./hello.js"
import {name,add} from "./user.js"

abc.sum(1,2);
console.log(name);
add(1,3);
```

# 三、Vue

MVVM思想

- M：model 包括数据和一些基本操作
- V：view 视图，页面渲染结果
- VM：View-model，模型与视图间的双向操作（无需开发人员干涉）

视图和数据通过VM绑定起来，model里有变化会自动地通过Directives填写到视view中，视图表单中添加了内容也会自动地通过DOM Listeners保存到模型中。

教程：https://cn.vuejs.org/v2/guide/

```


view     ViewModel  Model
只要我们Model发生了改变，View上自然就会表现出来。
当用户修改了View，Model中的数据也会跟着改变。
把开发人员从琐的D0M操作中解放出来，那关注点放在如何操作Model上。
```



#### 安装：

- 下载js并用 `<script>` 标签引入：`<script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script>`
- 或者在VScode控制台使用npm install vue导入。步骤分为：
  - 先`npm init -y`初始化项目，生成了一个`package.json`文件，说明他是一个npm管理的项目
    - 类似于maven的pom.xml
  - `npm install vue`，安装后在项目`node_modules`里有vue
    - 类似maven install拉取远程到本地

> npm install node-sass报错问题的．只要我们Model发生了改变，View上自然就会表现出来。
> 当用户修改了View，Model中的数据也厶跟着改变。
> 把开发人员从琐的D0M操作中解放出来，扌巴关/在如何操作Model上。解决：[https://blog.csdn.net/hancoder/article/details/113821646](https://blog.csdn.net/hancoder/article/details/113821646)

> `node_modules`目录的管理：
>
> 因为要提交给git push等原因，并不想把`node_modules`也提交上传，文件太大。
>
> 方式1是在`.gitignore`文件中指定不上传的文件夹
>
> 方式2是npm install安装包的时候指定安装的目录。`npm install --prefix path_to_node_modules放置的目录 `
>
> 方式3是在`package.json`中指定`export NODE_PATH='yourdir'/node_modules`
>
> 最终选择了方式1，别的方式不方便读者

使用：

- new Vue
- 在dom中{{name}}代表从模型中放到view中
- v-model实现双向绑定

index.html

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>

    <div id="app">
        <input type="text" v-model="num">
        v-model实现双向绑定。此处代表输入框和vue里的data绑定
        
        <button v-on:click="num++">点赞</button>
        v-on:click绑定事件，实现自增。
        
        <button v-on:click="cancel">取消</button>
        回调自定义的方法。 此时字符串里代表的函数
        
        <h1> {{name}} ,非常帅，有{{num}}个人为他点赞{{hello()}}</h1>
        先从vue中拿到值填充到dom，input再改变num值，vue实例更新，然后此处也更新
    </div>

    <!-- 导入依赖 -->
    <script src="./node_modules/vue/dist/vue.js"></script>

    <script>
        //1、vue声明式渲染
        let vm = new Vue({ //生成vue对象
            el: "#app",//绑定元素 div id="app" // 可以指定恰标签，但是不可以指定body标签
            data: {  //封装数据
                name: "张三",  // 也可以使用{} //表单中可以取出
                num: 1
            },
            methods:{  //封装方法
                cancel(){
                    this.num -- ;
                },
                hello(){
                    return "1"
                }
            }
        });
        // 还可以在html控制台vm.name

        //2、双向绑定,模型变化，视图变化。反之亦然。
        //3、事件处理

        //v-xx：指令

        //1、创建vue实例，关联页面的模板，将自己的数据（data）渲染到关联的模板，响应式的
        //2、指令来简化对dom的一些操作。
        //3、声明方法来做更复杂的操作。methods里面可以封装方法。

    </script>
</body>

</html>
```

在VSCode中安装vue 2 snippets语法提示插件，在谷歌浏览器中安装vue-devtool



##### 1、v-text、v-html、v-ref

这两个可以使用data数据。而`<div>123{{}}</div>`这种写法叫插值表达式，可以计算，可以取值，可以调用函数
这里还介绍v-html    v-text区别
注意取的大多数不是请求域了，而是vue对象里的data

> 插值闪烁：
>
> 使用{{}}方式在网速较慢时会出现问题。在数据未加载完成时，页面会显示出原始的`{{}}`，
> 加载完毕后才显示正确数据，我们称为插值闪烁。
> 我们将网速调慢一些，然后刷新页面，试试看刚才的案例



```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>
   
    <div id="app">
        {{msg}}  {{1+1}}  {{hello()}} 前面的内容如果网速慢的话会先显示括号，然后才替换成数据。
        v-html 和v-text能解决这个问题
        <br/>
        
        用v-html取内容
        <span v-html="msg"></span>
        
        <br/>
        原样显示
        <span v-text="msg"></span>  
    </div>
   
    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>
        new Vue({
            el:"#app",
            data:{
                msg:"<h1>Hello</h1>",
                link:"http://www.baidu.com"
            },
            methods:{
                hello(){
                    return "World"
                }
            }
        })
    </>
    
</body>
</html>
```

##### 2、**单向绑定v-bind:**

问题：花括号只能写在标签体内（`<div  标签内> 标签体 </div>`），不能用在标签内。

插值表达式只能用在标签体里，如果我们这么用`<a href="{{}}">`是不起作用的，所以需要 `<a v-bind:href="link">跳转</a>`这种用法

解决：用`v-bind:`，简写为`:`。表示把model绑定到view。可以设置src、title、class等

>  在浏览器里`vm.link="www.baidu.com"`，此处vue数据改了，dom里跳转链接也改了
>
>  class里有哪些内容可以通过vue数据的bool值添加删除，而在style中代表的是`k:v`值。
>
>  也可以把`v-bind:`简写成`:`
>
>  {{}}必须有返回值

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
</head>
<body>

    <!-- 给html标签的属性绑定 -->
    <div id="app"> 

        <a v-bind:href="link">跳转</a>

        <!-- class,style  {class名：vue值}-->
        <span v-bind:class="{active:isActive,'text-danger':hasError}"
          :style="{color: color1,fontSize: size}">你好</span>

    </div>

    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>
        let vm = new Vue({
            el:"#app",
            data:{
                link: "http://www.baidu.com",
                isActive:true,
                hasError:true,
                color1:'red',
                size:'36px'
            }
        })
    </script>

</body>
</html>
```

Vue.js 为 v-on 提供了事件修饰符来处理 DOM 事件细节，如：event.preventDefault() 或 event.stopPropagation()。

Vue.js 通过由点 **.** 表示的指令后缀来调用修饰符。

- `.stop` - 阻止冒泡
- `.prevent` - 阻止默认事件
- `.capture` - 阻止捕获
- `.self` - 只监听触发该元素的事件
- `.once` - 只触发一次
- `.left` - 左键事件
- `.right` - 右键事件
- `.middle` - 中间滚轮事件

```vue
<!-- 同上 -->
<input v-on:keyup.enter="submit">
<!-- 缩写语法 -->
<input @keyup.enter="submit">
```

##### 样式绑定v-bind:class

class 与 style 是 HTML 元素的属性，用于设置元素的样式，我们可以用 v-bind 来设置样式属性。

Vue.js v-bind 在处理 class 和 style 时， 专门增强了它。表达式的结果类型除了字符串之外，还可以是对象或数组。

```vue
实例中将 isActive 设置为 true 显示了一个绿色的 div 块，如果设置为 false 则不显示：
<div v-bind:class="{ 'active': isActive }"></div>


text-danger 类背景颜色覆盖了 active 类的背景色：
<div class="static"
     v-bind:class="{ 'active' : isActive, 'text-danger' : hasError }">
</div>
```

下面使用一个属性代表了若干class

```js
<div id="app">
    <div v-bind:class="classObject"></div>
</div>

new Vue({
    el: '#app',
    data: {
        isActive: true,
        error: {
            value: true,
            type: 'fatal'
        }
    },
    computed: {
        classObject: function () {
            return {
                base: true,
                active: this.isActive && !this.error.value,
                'text-danger': this.error.value && this.error.type === 'fatal',
            }
        }
    }
})
```

##### 3、**双向绑定v-model**

v-bind只能从model到view。v-model能从view到model

```html
<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>
<body>

    <!-- 表单项，自定义组件 -->
    <div id="app">

        精通的语言：如果是多选框，那么会把每个value值赋值给vue数据
            <input type="checkbox" v-model="language" value="Java"> java<br/>
            <input type="checkbox" v-model="language" value="PHP"> PHP<br/>
            <input type="checkbox" v-model="language" value="Python"> Python<br/>
        选中了 {{language.join(",")}}
    </div>
    
    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>
        let vm = new Vue({
            el:"#app",
            data:{
                language: []
            }
        })
    </script>

</body>
</html>
```

##### 4、v-on事件

事件监听可以使用 v-on 指令

`v-on:事件类型="方法"`  ，可以简写成`@事件类型="方法"`



>   事件冒泡：大小div都有单机事件，点了内部div相当于外部div也点击到了。
>
>   如果不想点击内部div冒泡到外部div，可以使用.prevent阻止事件冒泡
>
>   用法是`v-on:事件类型.事件修饰符="方法"`
>
>   还可以绑定按键修饰符
>
>   v-on:keyup.up="num+=2" @keyup.down="num-=2" @click.ctrl="num=10"

按键修饰符

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <div id="app">
                
        <!--事件中直接写js片段-->
        <button v-on:click="num++">点赞</button>
        <!--事件指定一个回调函数，必须是Vue实例中定义的函数-->
        <button @click="cancel">取消</button>
        <!--  -->
        <h1>有{{num}}个赞</h1>


        <!-- 事件修饰符 -->
        <div style="border: 1px solid red;padding: 20px;" v-on:click.once="hello">
            大div
            <div style="border: 1px solid blue;padding: 20px;" @click.stop="hello">
                小div <br />
                <a href="http://www.baidu.com" @click.prevent.stop="hello">去百度</a>
            </div>
        </div>



        <!-- 按键修饰符： -->
        <input type="text" v-model="num" v-on:keyup.up="num+=2" @keyup.down="num-=2" @click.ctrl="num=10"><br />

        提示：

    </div>
    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>
        new Vue({
            el:"#app",
            data:{
                num: 1
            },
            methods:{
                cancel(){
                    this.num--;
                },
                hello(){
                    alert("点击了")
                }
            }
        })
    </script>
</body>

</html>
```

##### 5、v-for遍历

可以遍历   数组[]  字典{}  。对于字典`<li v-for="(value, key, index) in object">`

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>

    <div id="app">
        <ul>
            <!-- 4、遍历的时候都加上:key来区分不同数据，提高vue渲染效率 -->
            <li v-for="(user,index) in users" :key="user.name" v-if="user.gender == '女'">
                <!-- 1、显示user信息：v-for="item in items" -->
               当前索引：{{index}} ==> {{user.name}}  ==>   
                  {{user.gender}} ==>{{user.age}} <br>
                <!-- 2、获取数组下标：v-for="(item,index) in items" -->
                <!-- 3、遍历对象：
                        v-for="value in object"
                        v-for="(value,key) in object"
                        v-for="(value,key,index) in object" 
                -->
                对象信息：
                <span v-for="(v,k,i) in user">{{k}}=={{v}}=={{i}}；</span>
                <!-- 4、遍历的时候都加上:key来区分不同数据，提高vue渲染效率 -->
            </li>

            
        </ul>

        <ul>
            <li v-for="(num,index) in nums" :key="index"></li>
        </ul>
    </div>
    <script src="../node_modules/vue/dist/vue.js"></script>
    <script>         
        let app = new Vue({
            el: "#app",
            data: {
                users: [
                { name: '柳岩', gender: '女', age: 21 },
                { name: '张三', gender: '男', age: 18 },
                { name: '范冰冰', gender: '女', age: 24 },
                { name: '刘亦菲', gender: '女', age: 18 },
                { name: '古力娜扎', gender: '女', age: 25 }
                ],
                nums: [1,2,3,4,4]
            },
        })
    </script>
</body>

</html>
```



```js
new Vue({
  el: '#blog-post-demo',
  data: {
    posts: [
      { id: 1, title: 'My journey with Vue' },
      { id: 2, title: 'Blogging with Vue' },
      { id: 3, title: 'Why Vue is so fun' }
    ]
  }
})

<blog-post
  v-for="post in posts"
  v-bind:key="post.id"
  v-bind:title="post.title"
></blog-post>
```

```js
<blog-post
  v-for="post in posts"
  v-bind:key="post.id"
  v-bind:post="post">
</blog-post>

Vue.component('blog-post', {
  props: ['post'],
  template: `
    <div class="blog-post">
      <h3>{{ post.title }}</h3>
      <div v-html="post.content"></div>
    </div>
  `
})
```

##### 6、v-if和v-show

在vue实例的data指定一个bool变量，然后v-show赋值即可。

show里的字符串也可以比较

if是根据表达式的真假，切换元素的显示和隐藏（操作dom元素）

区别：show的标签F12一直都在，if的标签会移除，

if操作dom树对性能消耗大

v-if和v-show.html

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <!-- 
        v-if，顾名思义，条件判断。当得到结果为true时，所在的元素才会被渲染。
        v-show，当得到结果为true时，所在的元素才会被显示。 
    -->
    <div id="app">
        <button v-on:click="show = !show">点我呀</button>
        <!-- 1、使用v-if显示 -->
        <h1 v-if="show">if=看到我....</h1>
        <!-- 2、使用v-show显示 -->
        <h1 v-show="show">show=看到我</h1>
    </div>

    <script src="../node_modules/vue/dist/vue.js"></script>
        
    <script>
        let app = new Vue({
            el: "#app",
            data: {
                show: true
            }
        })
    </script>

</body>

</html>
```

##### 7、v-else和v-else-if

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <div id="app">
        <button v-on:click="random=Math.random()">点我呀</button>
        <span>{{random}}</span>

        <h1 v-if="random>=0.75">
            看到我啦? &gt;= 0.75
        </h1>
        <h1 v-else-if="random>=0.5">
            看到我啦? &gt;= 0.5
        </h1>
        <h1 v-else-if="random>=0.2">
            看到我啦? &gt;= 0.2
        </h1>
        <h1 v-else>
            看到我啦? &lt; 0.2
        </h1>

    </div>


    <script src="../node_modules/vue/dist/vue.js"></script>
        
    <script>         
        let app = new Vue({
            el: "#app",
            data: { random: 1 }
        })     
    </script>
</body>

</html>
```

##### 缩写@

Vue.js 为两个最为常用的指令提供了特别的缩写：

v-bind 缩写

```vue
<!-- 完整语法 -->
<a v-bind:href="url"></a>
<!-- 缩写 -->
<a :href="url"></a>
```

v-on 缩写

```vue
<!-- 完整语法 -->
<a v-on:click="doSomething"></a>
<!-- 缩写 -->
<a @click="doSomething"></a>
```

### 2、计算属性和侦听器

##### 计算属性computed

什么是计算属性：**属性**不是具体值，而是通过一个函数计算出来的，随时变化

```vue
<div id="app">
    <p>原始字符串: {{ message }}</p>
    <p>计算后反转字符串: {{ reversedMessage }}</p>
</div>

<script>
    var vm = new Vue({
        el: '#app',
        data: {
            message: 'Runoob!'
        },
        computed: {
            // 计算属性的 getter
            reversedMessage: function () {
                // `this` 指向 vm 实例
                return this.message.split('').reverse().join('')
            }
        }
    })
</script>
```

这里可以比较一下计算和method之间的区别

```js
computed: {
    // 计算属性的 getter
    reversedMessage: function () {
        // `this` 指向 vm 实例
        return this.message.split('').reverse().join('')
    }
},

methods: {
    reversedMessage2: function () {
        return this.message.split('').reverse().join('')
    }
}
```



##### 监听$watch

监听属性 watch，我们可以通过 watch 来响应数据的变化。

以下实例通过使用 watch 实现计数器：

```vue
<div id = "app">
    <p style = "font-size:25px;">计数器: {{ counter }}</p>
    <button @click = "counter++" style = "font-size:25px;">点我</button>
</div>
<script type = "text/javascript">
    var vm = new Vue({
        el: '#app',
        data: {
            counter: 1
        }
    });
    vm.$watch('counter', function(nval, oval) { // new old
        alert('计数器值的变化 :' + oval + ' 变为 ' + nval + '!');
    });
</script>
```



```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <div id="app">
        <!-- 某些结果是基于之前数据实时计算出来的，我们可以利用计算属性。来完成 -->
        <ul>
            <li>西游记； 价格：{{xyjPrice}}，数量：<input type="number" v-model="xyjNum"> </li>
            <li>水浒传； 价格：{{shzPrice}}，数量：<input type="number" v-model="shzNum"> </li>
            <li>总价：{{totalPrice}}</li>
            {{msg}}
        </ul>
    </div>
    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>
        //watch可以让我们监控一个值的变化。从而做出相应的反应。
        new Vue({
            el: "#app",
            data: {
                xyjPrice: 99.98,
                shzPrice: 98.00,
                xyjNum: 1,
                shzNum: 1,
                msg: ""
            },
            computed: {
                totalPrice(){
                    return this.xyjPrice*this.xyjNum + this.shzPrice*this.shzNum
                }
            },
            watch: {
                xyjNum: function(newVal,oldVal){
                    if(newVal>=3){
                        this.msg = "库存超出限制";
                        this.xyjNum = 3
                    }else{
                        this.msg = "";
                    }
                }
            },
        })
    </script>

</body>

</html>
```

##### 过滤器filter

定义filter组件后，管道符后面跟具体过滤器`{{user.gender | gFilter}}`

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <!-- 过滤器常用来处理文本格式化的操作。过滤器可以用在两个地方：双花括号插值和 v-bind 表达式 -->
    <div id="app">
        <ul>
            <li v-for="user in userList">
                {{user.id}} ==> {{user.name}} ==> {{user.gender == 1?"男":"女"}} ==>
                {{user.gender | genderFilter}} ==> {{user.gender | gFilter}}
            </li>
        </ul>
    </div>
    <script src="../node_modules/vue/dist/vue.js"></script>

    <script>

        // 全局过滤器
        Vue.filter("gFilter", function (val) {
            if (val == 1) {
                return "男~~~";
            } else {
                return "女~~~";
            }
        })

        let vm = new Vue({
            el: "#app",
            data: {
                userList: [
                    { id: 1, name: 'jacky', gender: 1 },
                    { id: 2, name: 'peter', gender: 0 }
                ]
            },
            filters: { // 局部过滤器，只可以在当前vue实例中使用
                genderFilter(val) {
                    if (val == 1) {
                        return "男";
                    } else {
                        return "女";
                    }
                }
            }
        })
    </script>
</body>

</html>
```

### 3、表单、复选框

https://www.runoob.com/vue2/vue-forms.html

##### 1) 多选checkbox

虽然v-model指定了同个值，但是会收集成数组。演示：https://www.runoob.com/try/try.php?filename=vue2-form2

![img](https://i0.hdslb.com/bfs/album/aee01a97871288d38b3fdb916e06297d23cfe294.jpg)

```VUE
<body>
    <div id="app">
        <p>单个复选框：</p>
        <input type="checkbox" id="checkbox" v-model="checked">
        <label for="checkbox">{{ checked }}</label>

        <p>多个复选框：</p>
        <input type="checkbox" id="runoob" value="Runoob" v-model="checkedNames">
        <label for="runoob">Runoob</label>
        <input type="checkbox" id="google" value="Google" v-model="checkedNames">
        <label for="google">Google</label>
        <input type="checkbox" id="taobao" value="Taobao" v-model="checkedNames">
        <label for="taobao">taobao</label>
        <br>
        <span>选择的值为: {{ checkedNames }}</span>
    </div>

    <script>
        new Vue({
            el: '#app',
            data: {
                checked : false,
                checkedNames: []
            }
        })
    </script>
</body>
```

##### 2) 单选radio

```vue
<div id="app">
    <input type="radio" id="runoob" value="Runoob" v-model="picked">
    <label for="runoob">Runoob</label>
    <br>
    <input type="radio" id="google" value="Google" v-model="picked">
    <label for="google">Google</label>
    <br>
    <span>选中值为: {{ picked }}</span>
</div>

<script>
    new Vue({
        el: '#app',
        data: {
            picked : 'Runoob'
        }
    })
</script>
```

##### 3) 下拉select

```vue
<div id="app">
    <select v-model="selected" name="fruit">
        <option value="">选择一个网站</option>
        <option value="www.runoob.com">Runoob</option>
        <option value="www.google.com">Google</option>
    </select>

    <div id="output">
        选择的网站是: {{selected}}
    </div>
</div>

<script>
    new Vue({
        el: '#app',
        data: {
            selected: '' 
        }
    })
</script>
```

##### 4) 修饰符

**.lazy**

在默认情况下， v-model 在 input 事件中同步输入框的值与数据，但你可以添加一个修饰符 lazy ，从而转变为在 change 事件中同步：

```
<!-- 在 "change" 而不是 "input" 事件中更新 -->
<input v-model.lazy="msg" >
```

**.number**

如果想自动将用户的输入值转为 Number 类型（如果原值的转换结果为 NaN 则返回原值），可以添加一个修饰符 number 给 v-model 来处理输入值：

```
<input v-model.number="age" type="number">
```

这通常很有用，因为在 type="number" 时 HTML 中输入的值也总是会返回字符串类型。

**.trim**

如果要自动过滤用户输入的首尾空格，可以添加 trim 修饰符到 v-model 上过滤输入：

```
<input v-model.trim="msg">
```

### 4、==组件化==

在大型应用开发的时候，页面可以划分成很多部分。往往不同的页面，也会有相同的部分。例如可能会有相同的**头部导航**。

但是如果每个页面都自开发，这无疑增加了我们开发的成本。所以我们会把页面的不同分拆分成立的组件，然后在不同页面就可以共享这些组件，避免重复开发。

在vue里，所有的vue实例都是组件

- ==组件其实也是一个vue实例==，因此它在定义时也会接收：data、methods、生命周期函数等
- 不同的是组件不会与页面的元素绑定（所以不写el），否则就无法复用了，因此没有`el`属性。
- 但是组件渲染需要html模板，所以增加了`template`属性，值就是HTML模板
- data必须是一个**函数**，不再是一个对象。
- 全局组件定义完毕，任何vue实例都可以直接在HTML中通过组件名称来使用组件了

下面的html内容是在说：

- `<div id="app">`这个块跟下面的new Vue对象绑定
- vue对象里指定了components属性，是指：标签里有子标签，`'button-counter': buttonCounter`代表有个组件叫buttonCounter，可以填充标签`<button-counter>`

#### vue组件的使用

参考：https://cn.vuejs.org/v2/guide/components.html

https://www.runoob.com/vue2/vue-component.html

#### 1) 全局组件

参考：https://cn.vuejs.org/v2/guide/components-registration.html

```JS
// 定义一个名为 button-counter 的新组件
Vue.component('button-counter', {
  data: function () { // data是函数
    return {
      count: 0
    }
  },
  template: '<button v-on:click="count++">You clicked me {{ count }} times.</button>' // 渲染的模板
})

// 上面代表组件名为button-counter   
会拿 <button v-on:click="count++">You clicked me {{ count }} times.</button> 替代
```

全局注册的组件可以用在其被注册之后的任何 (通过 `new Vue`) 新创建的 Vue 根实例，也包括其组件树中的所有**子组件**的模板中。

>  这里还应该注解的是data那个属性，在组件中需要以函数的方式写



#### 2) 局部组件

注意：==局部注册的组件在其子组件中不可用==

```VUE
<div id="app">
    <runoob></runoob>
</div>

<script>
    
    // 声明一个组件
    // 这个东西可以在当前文件中写，也可以import
    var Child = {
        template: '<h1>自定义组件!</h1>'
        // 不定义el
        // 可以定义data函数
    }

    // 创建根实例(它是使用组件的)
    new Vue({
        el: '#app',
        components: {
            // <runoob> 将只在父模板可用  别的组件如果想使用其他局部组件，必须指定
            'runoob': Child //局部组件在div中感应不到，所以要绑定一下，原来是就是说碰到runoob标签就替换为template内容
        }
    })
</script>
```

> 关于import和export可以看之前ES处的内容，但是在一个vue文件里导出，在另外vue组件里导入使用而已
>
> 带default代表就是导出了，你可以接收时候可以自己命名别名。而不加default导出时，当前文件导入的时候得按原名来

```js
var ComponentA = { /* ... */ }

var ComponentB = {
  components: {
    'component-a': ComponentA
  },
  // ...
}
```

或者如果你通过 Babel 和 webpack 使用 ES2015 模块，那么代码看起来更像：

```js
import ComponentA from './ComponentA.vue'

export default {
  components: {
    ComponentA
  },
  // ...
}
```

注意在 ES2015+ 中，在对象中放一个类似 `ComponentA` 的变量名其实是 `ComponentA: ComponentA` 的缩写，即这个变量名同时是：

- 用在模板中的自定义元素的名称
- 包含了这个组件选项的变量名



##### prop

prop 是子组件用来接受父组件传递过来的数据的一个自定义属性。

父组件的数据需要通过 props 把数据传给子组件，子组件需要显式地用 props 选项声明 "prop"：

```vue
<div id="app">
    <child message="hello!"></child>
</div>

<script>
    // 注册全局组件
    Vue.component('child', {
        // 声明 props
        props: ['message'],
        // 同样也可以在 vm 实例中像 "this.message" 这样使用
        template: '<span>{{ message }}</span>'
    })
    
    // 定义vue对象
    new Vue({
        el: '#app'
    })
</script>
```



##### 组件hello world

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
</head>

<body>

    <div id="app">
        <button v-on:click="count++">我被点击了 {{count}} 次</button>

        每个对象都是独立统计的
        <counter></counter>
        <counter></counter>
        <counter></counter>
        <counter></counter>
        <counter></counter>

        <button-counter></button-counter>
    </div>
    <script src="../node_modules/vue/dist/vue.js"></script>


    <script>
        //1、全局声明注册一个组件 // counter标签，代表button
        // 把页面中<counter>标签替换为指定的template，而template中的数据用data填充
        Vue.component("counter", {
            template: `<button v-on:click="count++">我被点击了 {{count}} 次</button>`,
            data() {// 如果 Vue 没有这条规则，点击一个按钮就可能会像如下代码一样影响到其它所有实例：
                return {
                    count: 1 // 数据
                }
            }
        });

        //2、局部声明一个组件
        const buttonCounter = {
            template: `<button v-on:click="count++">我被点击了 {{count}} 次~~~</button>`,
            data() {
                return {
                    count: 1
                }
            }
        };

        new Vue({
            el: "#app",
            data: {
                count: 1
            },
            components: { // 要用的组件
                'button-counter': buttonCounter
            }
        })
    </script>
</body>

</html>
```



#### **父子组件**

我自己倾向于说父标签是子组件，子标签是父组件，因为父标签的vue定义时有个template属性指定了模板，也就是继承了父组件(子标签)

> 父子组件是什么：也没有个确定的答案https://www.zhihu.com/question/55892803
>
> ![img](https://pic1.zhimg.com/80/v2-691f97e24a12aba5e5f3ef77a854e57a_720w.jpg?source=1940ef5c)
>
> 父实例(也就是#app)里的红色部分为父组件，紫色部分为父组件模板
>
> 注册的全局组件时，红色部分为子组件，紫色部分为子组件模板。
>
> - 在vue中，父子组件的关系可以总结为`prop`向下传递，`事件`向上传递。父组件通过`prop`给子组件下发数据，子组件通过`事件`给父组件发送信息。

### 5、生命周期钩子函数

每个vue实例在被创建时都要经过一系列的初始化过程：创建实例，装载模板、渲染模板等等。vue为生命周期中的每个状态都设置了钩子函数（监听函）。每当vue实列处于不同的生命周期时，对应的函数就会被触发调用。



<img src="https://cn.vuejs.org/images/lifecycle.png" alt="Vue 实例生命周期" style="zoom: 33%;" />

```html
<!DOCTYPE html>
<html lang="en">

<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta http-equiv="X-UA-Compatible" content="ie=edge">
    <title>Document</title>
</head>

<body>
    <div id="app">
        <span id="num">{{num}}</span>
        <button @click="num++">赞！</button>
        <h2>{{name}}，有{{num}}个人点赞</h2>
    </div>

    <script src="../node_modules/vue/dist/vue.js"></script>
    
    <script>
        let app = new Vue({
            el: "#app",
            data: {
                name: "张三",
                num: 100
            },
            methods: {
                show() {
                    return this.name;
                },
                add() {
                    this.num++;
                }
            },
            beforeCreate() {
                console.log("=========beforeCreate=============");
                console.log("数据模型未加载：" + this.name, this.num);
                console.log("方法未加载：" + this.show());
                console.log("html模板未加载：" + document.getElementById("num"));
            },
            created: function () {
                console.log("=========created=============");
                console.log("数据模型已加载：" + this.name, this.num);
                console.log("方法已加载：" + this.show());
                console.log("html模板已加载：" + document.getElementById("num"));
                console.log("html模板未渲染：" + document.getElementById("num").innerText);
            },
            beforeMount() {
                console.log("=========beforeMount=============");
                console.log("html模板未渲染：" + document.getElementById("num").innerText);
            },
            mounted() {
                console.log("=========mounted=============");
                console.log("html模板已渲染：" + document.getElementById("num").innerText);
            },
            beforeUpdate() {
                console.log("=========beforeUpdate=============");
                console.log("数据模型已更新：" + this.num);
                console.log("html模板未更新：" + document.getElementById("num").innerText);
            },
            updated() {
                console.log("=========updated=============");
                console.log("数据模型已更新：" + this.num);
                console.log("html模板已更新：" + document.getElementById("num").innerText);
            }
        });
    </script>
</body>

</html>
```

### 6、路由

Vue.js 路由允许我们通过不同的 URL 访问不同的内容。

通过 Vue.js 可以实现多视图的单页Web应用（single page web application，SPA）。

> ````sh
> cnpm install vue-router
> ````

Vue.js + vue-router 可以很简单的实现单页应用。

`<router-link>`是一个组件，该组件用于设置一个导航链接，切换不同 HTML 内容。 **to** 属性为目标地址， 即要显示的内容。

以下实例中我们将 vue-router 加进来，然后配置组件和路由映射，再告诉 vue-router 在哪里渲染它们

```VUE
<script src="https://unpkg.com/vue/dist/vue.js"></script>
<script src="https://unpkg.com/vue-router/dist/vue-router.js"></script>
 
<div id="app">
  <h1>Hello App!</h1>
  <p>
    <!-- 使用 router-link 组件来导航. -->
    <!-- 通过传入 `to` 属性指定链接. -->
    <!-- <router-link> 默认会被渲染成一个 `<a>` 标签 -->
    <router-link to="/foo">Go to Foo</router-link>
    <router-link to="/bar">Go to Bar</router-link>
  </p>
  <!-- 路由出口 -->
  <!-- 路由匹配到的组件将渲染在这里 -->
  <router-view></router-view>
</div>
```

```JS
// 0. 如果使用模块化机制编程，导入 Vue 和 VueRouter，要调用 Vue.use(VueRouter)

// 1. 定义（路由）组件。
// 可以从其他文件 import 进来
const Foo = { template: '<div>foo</div>' }
const Bar = { template: '<div>bar</div>' }

// 2. 定义路由
// 每个路由应该映射一个组件。 其中"component" 可以是
// 通过 Vue.extend() 创建的组件构造器，
// 或者，只是一个组件配置对象。
// 我们晚点再讨论嵌套路由。
const routes = [
    { path: '/foo', component: Foo },
    { path: '/bar', component: Bar }
]

// 3. 创建 router 实例，然后传 `routes` 配置
// 你还可以传别的配置参数, 不过先这么简单着吧。
const router = new VueRouter({
    routes // （缩写）相当于 routes: routes
})

// 4. 创建和挂载根实例。
// 记得要通过 router 配置参数注入路由，
// 从而让整个应用都有路由功能
const app = new Vue({
    router
}).$mount('#app')

// 现在，应用已经启动了！
```





## 四、vue项目

观察renren-fast-vue，只有一个index.html，其他都是vue文件夹

```
index.html
static
config
src/APP.vue
src/views/main.vue
src/views/main-navbar.vue
```

```bash
（这个是工具，与具体项目无关）

# 安装webpack
npm install webpack -g

# 安装vue脚手架
npm install -g @vue/cli-init  
```

```sh
# 在项目文件夹里执行
# 初始化vue项目
vue init webpack vue-demo


# 运行，访问8080端口
npm run dev 
```





- 初始化vue项目：vue init webpack appTestName项目名
  - 我直接在IDEA的项目下（不用创建文件夹）执行`vue init webpack vue-demo`
  - 报错的话添加vue的环境变量：
    - npm config get prefix  找到npm
    - 添加环境变量 C:\Users\HAN\AppData\Roaming\npm\node_modules\@vue\cli-init\node_modules\.bin
- 一个劲回车，ESlint及后面的都选择no
- npm run dev 运行，访问8080端口

#### vue项目目录结构：

| 目录/文件      | 说明                                                         |
| :------------- | :----------------------------------------------------------- |
| build          | 项目构建(webpack)相关代码                                    |
| `config`       | 配置目录，包括端口号等。我们初学可以使用默认的。             |
| `node_modules` | npm 加载的项目依赖模块                                       |
| src            | 这里是我们要开发的目录，基本上要做的事情都在这个目录里。里面包含了几个目录及文件：<br />- `assets`: 放置一些图片，如logo等。<br />- `components`: 目录里面放了一个组件文件，可以不用。<br />- `App.vue`: 项目入口文件，我们也可以直接将组件写这里，而不使用 components 目录。<br />- `main.js`: 项目的核心文件。 |
| static         | 静态资源目录，如图片、字体等。                               |
| test           | 初始测试目录，可删除                                         |
| .xxxx文件      | 这些是一些配置文件，包括语法配置，git配置等。                |
| `index.html`   | 首页入口文件，你可以添加一些 meta 信息或统计代码啥的。       |
| `package.json` | 项目配置文件。                                               |
| README.md      | 项目的说明文档，markdown 格式                                |

> cmd属性，去掉快速编辑

#### ==关系结构==

##### 1) index.html

index.html只有一个`<div id="app">`

```html
<script>document.write('<script src="./config/index.js?t=' + new Date().getTime() + '"><\/script>');</script>

<body>
  <div id="app"></div>
</body>
```

##### 2) main.js

main.js中import，并且new Vue

```js
import Vue from 'vue'
import App from '@/App'
import router from '@/router'                 // api: https://github.com/vuejs/vue-router
import store from '@/store' 

new Vue({
  el: '#app', // index.html中的div
  router,
  store,
  template: '<App/>',
  components: { App }
})

在 Vue 构造器中有一个el 参数，它是 DOM 元素中的 id。
这意味着我们接下来的改动全部在以上指定的 div 内，div 外部不受影响。
```

##### 3) App.vue

```vue
<template>
  <transition name="fade">
    <router-view></router-view>
  </transition>
</template>

<script>
    // 导出vue实例。相当于合并了以前的new vue和export导出组件
  export default {
  }
</script>
```

- 其中的路由视图标签是根据url要决定访问的vue
- 在main.js中提及了是使用的./router

component是导入了组件，template是使用了组件

attrgroup.vue      (为了使结构清晰，删除很多代码)

```VUE
<template>
  <el-row :gutter="20">
    <el-col :span="6">
      <category @tree-node-click="treenodeclick"></category>
    </el-col>
      
    <el-col :span="18">
              <div class="mod-config">
        <el-form :inline="true" :model="dataForm" @keyup.enter.native="getDataList()">
          <el-form-item>
            <el-button @click="getDataList()">查询</el-button>
          </el-form-item>
        </el-form>
        <el-table :data="dataList">
          <el-table-column type="selection"。。。></el-table-column>
          <el-table-column>。。。 </el-table-column>
        </el-table>

        <el-pagination 。。。></el-pagination>

        <!-- 弹窗, 新增 / 修改 -->
        <add-or-update v-if="addOrUpdateVisible" ref="addOrUpdate" @refreshDataList="getDataList"></add-or-update>

      </div>
    </el-col>
  </el-row>
</template>

<script>
import Category from "../common/category";
import AddOrUpdate from "./attrgroup-add-or-update";
import RelationUpdate from "./attr-group-relation";
export default {
  //import引入的组件需要注入到对象中才能使用。组件名:自定义的名字，一致可以省略
  components: { Category, AddOrUpdate, RelationUpdate },
  props: {},
  data() {
    return {
      catId: 0,
      dataForm: {
        key: ""
      },
      dataList: [],
      pageIndex: 1,
      pageSize: 10,
      totalPage: 0,
      dataListLoading: false,
      dataListSelections: [],
      addOrUpdateVisible: false,
      relationVisible: false
    };
  },
  activated() {
    this.getDataList();
  },
  methods: {
```



##### vue.json

文件--首选项--用户代码片段--新建代码片段--取名vue.json

```json
// https://www.cnblogs.com/songjilong/p/12635448.html
{
    "Print to console": {
        "prefix": "vue",
        "body": [
            "<!-- $1 -->",
            "<template>",
            "<div class='$2'>$5</div>",
            "</template>",
            "",
            "<script>",
            "//这里可以导入其他文件（比如：组件，工具js，第三方插件js，json文件，图片文件等等）",
            "//例如：import 《组件名称》 from '《组件路径》';",
            "",
            "export default {",
            "//import引入的组件需要注入到对象中才能使用",
            "components: {},",
            "data() {",
            "//这里存放数据",
            "return {",
            "",
            "};",
            "},",
            "//监听属性 类似于data概念",
            "computed: {},",
            "//监控data中的数据变化",
            "watch: {},",
            "//方法集合",
            "methods: {",
            "",
            "},",
            "//生命周期 - 创建完成（可以访问当前this实例）",
            "created() {",
            "",
            "},",
            "//生命周期 - 挂载完成（可以访问DOM元素）",
            "mounted() {",
            "",
            "},",
            "beforeCreate() {}, //生命周期 - 创建之前",
            "beforeMount() {}, //生命周期 - 挂载之前",
            "beforeUpdate() {}, //生命周期 - 更新之前",
            "updated() {}, //生命周期 - 更新之后",
            "beforeDestroy() {}, //生命周期 - 销毁之前",
            "destroyed() {}, //生命周期 - 销毁完成",
            "activated() {}, //如果页面有keep-alive缓存功能，这个函数会触发",
            "}",
            "</script>",
            "<style scoped>",
            "//@import url($3); 引入公共css类",
            "$4",
            "</style>"
        ],
        "description": "生成vue模板"
    },
    "http-get请求": {
	"prefix": "httpget",
	"body": [
		"this.\\$http({",
		"url: this.\\$http.adornUrl(''),",
		"method: 'get',",
		"params: this.\\$http.adornParams({})",
		"}).then(({ data }) => {",
		"})"
	],
	"description": "httpGET请求"
    },
    "http-post请求": {
	"prefix": "httppost",
	"body": [
		"this.\\$http({",
		"url: this.\\$http.adornUrl(''),",
		"method: 'post',",
		"data: this.\\$http.adornData(data, false)",
		"}).then(({ data }) => { });" 
	],
	"description": "httpPOST请求"
    }
}
```



> new Vue中指定了元素，元素如何渲染看template，组件决定元素样式

## 五、element-ui

官网： https://element.eleme.cn/#/zh-CN/component/installation 

安装

```sh
# 直接npm安装，在项目中执行
npm i element-ui -S

# 或者引入样式
```



使用： 

在 main.js 中写入以下内容： 

```js
import ElementUI  from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';

// 让vue使用ElementUI组件
Vue.use(ElementUI);
```

然后.vue文件中写标签 





#### 怎么改

main.js中有

```js
new Vue({
  el: '#app', // 他是index.html中的div
  router,//router:router，是从上面的./router里来的 // 会替换<router-view/>
  components: { App },//App:App  属性名和属性值一样时可以简写 // 要用的组件
  template: '<App/>' // 导入的<template> <div id="app">
})
```

- el表示与index.html中的div关联，
- 在div里面要展示App，所以把App这个组件修改
- 在App.vue里说了export name 'App'是说把模板命名为App，所以我们替换他的模板

### **插槽**

由于插槽是一块模板，所以，对于任何一个组件，从模板种类的角度来分，其实都可以分为**非插槽模板**和**插槽模板**两大类。

- 非插槽模板指的是**html模板**，指的是‘div、span、ul、table’这些，非插槽模板的显示与隐藏以及怎样显示由插件自身控制；
- 插槽模板是slot，它是一个空壳子，因为它显示与隐藏以及最后用什么样的**html模板**显示由父组件（子标签）控制。**但是插槽显示的位置确由子组件自身决定，slot写在组件template的哪块，父组件传过来的模板将来就显示在哪块**。

下面插槽内容参考：https://segmentfault.com/a/1190000012996217

之前我们怎么写：定义组件之后使用

```html

<div id="app">
    <child content='<p>哈哈哈</p>'></child>
</div>

<script>
    Vue.component('child', {
        props: [
            'content'
        ],
        template: '<div ><p>hello</p><div v-html="this.content"></div></div>'
    })
    var vm = new Vue({
        el: '#app'
    })
</script>
```

上面代码叫做**通过属性传递**，缺点:会多出一个div，template 不好用，所以通过content传值，直接使用p标签是有问题的，必须包一个div。传递少还行，当传递很多的时候，代码会很难去阅读

```html

<div id="app">
    <child>
        <!-- 1.用子组件时插入一点内容 -->
        <h1>hahaha</h1>
    </child>
</div>
<script>
    Vue.component('child', {
        // 2.用<slot></slot>可以显示  <p>hahaha</p>
        // 所以使用插槽会更方便的向子组件插入dom元素
        template: `<div>
                <p>hello</p>
                <slot></slot>
                    </div>`
     })
     
    var vm = new Vue({
        el: '#app'
    })
</script>
```

slot有新的特性
1.定义默认值，定义在slot标签上
2.具名插槽

```html
<div id="app">
    <child>
        <div class="footer" slot="footer">footer</div>
    </child>
</div>
<script>
    Vue.component('child', {
        template: `<div>
                    <slot name="header"></slot>
                    <div class="content">content</div>
                    <slot name="footer"></slot>
                   </div>`
    })
    var vm = new Vue({
        el: '#app'
    })
</script>
```



#### 1）匿名插槽

首先是单个插槽、**单个插槽**是vue的官方叫法，但是其实也可以叫它**默认插槽**，或者与具名插槽相对，我们可以叫它匿名插槽。因为它不用设置name属性。

- 单个插槽可以放置在组件的任意位置，但是就像它的名字一样，一个组件中只能有一个该类插槽。
- 相对应的，具名插槽就可以有很多个，只要名字（name属性）不同就可以了。

下面通过一个例子来展示。

<img src="https://segmentfault.com/img/remote/1460000012996222?w=782&amp;h=342" alt="img" style="zoom:50%;" />

父组件：`</template>`标签表明是个模板

```vue
<template>
    <div class="father">
        <h3>这里是父组件</h3>
        <child>
            <div class="tmpl">
              <span>菜单1</span>
              <span>菜单2</span>
              <span>菜单3</span>
              <span>菜单4</span>
              <span>菜单5</span>
              <span>菜单6</span>
            </div>
        </child>
    </div>
</template>

<script>
  import Child from './Child.vue'
    
    export default {
        data: function () {
            return {
                msg: ''
            }
        },
        components:{
          'child': Child
        }
    }
</script>

<style scoped>
  .father{
    width:100%;
    background-color: #ccc;
    height: 650px;
  }
  .tmpl{
    display: flex;
    justify-content: space-around;
    flex-direction: row;
    width: 30%;
    margin: 0 auto;
  }
  .tmpl span{
    border:1px solid red;
    height:50px;
    line-height: 50px;
    padding: 10px;
  }
</style>
```

子组件：`<slot>`找一个默认的`</template>`

```vue
<template>
    <div class="child">
        <h3>这里是子组件</h3>
        <slot></slot>
    </div>
</template>

<script>
export default {
    data: function(){
        return {
            msg: ''
        }
    },
}
</script>

<style scoped>
  .child{
    background-color: #00bbff;
    width: 100%;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
</style>

```

什么意思呢？子组件里原来是有点内容的，不能全部用模板替换，这部分内容填充到父组件的`<slot>`里




```
注：所有demo都加了样式，以方便观察。其中，父组件以灰色背景填充，子组件都以浅蓝色填充。
```

#### 2）具名插槽

匿名插槽没有name属性，所以是匿名插槽，那么，插槽加了name属性，就变成了具名插槽。具名插槽可以在一个组件中出现N次。出现在不同的位置。下面的例子，就是一个有两个**具名插槽**和**单个插槽**的组件，这三个插槽被父组件用同一套css样式显示了出来，不同的是内容上略有区别。

父组件：

```VUE
<template>
  <div class="father">
    <h3>这里是父组件</h3>
    <child>
      <div class="tmpl" slot="up">
        <span>菜单1</span>
        <span>菜单2</span>
        <span>菜单3</span>
        <span>菜单4</span>
        <span>菜单5</span>
        <span>菜单6</span>
      </div>
      <div class="tmpl" slot="down">
        <span>菜单-1</span>
        <span>菜单-2</span>
        <span>菜单-3</span>
        <span>菜单-4</span>
        <span>菜单-5</span>
        <span>菜单-6</span>
      </div>
      <div class="tmpl">
        <span>菜单->1</span>
        <span>菜单->2</span>
        <span>菜单->3</span>
        <span>菜单->4</span>
        <span>菜单->5</span>
        <span>菜单->6</span>
      </div>
    </child>
  </div>
</template>

<script>
  import Child from './Child.vue'
  export default {
    data: function () {
      return {
        msg: ''
      }
    },
    components:{
      'child': Child
    }
  }
</script>

<style scoped>
  .father{
    width:100%;
    background-color: #ccc;
    height: 650px;
  }
  .tmpl{
    display: flex;
    justify-content: space-around;
    flex-direction: row;
    width: 30%;
    margin: 0 auto;
  }
  .tmpl span{
    border:1px solid red;
    height:50px;
    line-height: 50px;
    padding: 10px;
  }
</style>

```

子组件：

```VUE
<template>
  <div class="child">
    <slot name="up"></slot>

    <h3>这里是子组件</h3>

    <slot name="down"></slot>
    <slot></slot>
  </div>
</template>

<script>
  export default {
    data: function(){
      return {
        msg: ''
      }
    },
    computed: {
    },
    methods:{
    },
    components: {
    }
  }
</script>

<style scoped>
  .child{
    background-color: #00bbff;
    width: 100%;
    padding: 10px;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
</style>

```

显示结果如图：
![img](https://segmentfault.com/img/remote/1460000012996223?w=742&h=456)

可以看到，父组件通过html模板上的slot属性关联具名插槽。没有slot属性的html模板默认关联匿名插槽。

#### 3）作用域插槽 | 带数据的插槽

最后，就是我们的作用域插槽。这个稍微难理解一点。官方叫它作用域插槽，实际上，对比前面两种插槽，我们可以叫它带数据的插槽。什么意思呢，就是前面两种，都是在组件的template里面写

```HTML
匿名插槽
<slot></slot>
具名插槽
<slot name="up"></slot>
```

但是**作用域插槽要求，在slot上面绑定数据**。也就是你得写成大概下面这个样子。

```vue
<slot name="up" :data="data"></slot>
 export default {
    data: function(){
      return {
        data: ['zhangsan','lisi','wanwu','zhaoliu','tianqi','xiaoba']
      }
    },
}
```

我们前面说了，插槽最后显示不显示是看父组件有没有在child下面写模板，像下面那样。

```HTML
<child>
   html模板
</child>
```

写了，插槽就总得在浏览器上显示点东西，东西就是html该有的模样，没写，插槽就是空壳子，啥都没有。
OK，我们说有html模板的情况，就是父组件会往子组件插模板的情况，那到底插一套什么样的样式呢，这由父组件的html+css共同决定，但是这套样式里面的内容呢？

正因为作用域插槽绑定了一套数据，父组件可以拿来用。于是，情况就变成了这样：样式父组件说了算，但内容可以显示子组件插槽绑定的。

我们再来对比，作用域插槽和单个插槽和具名插槽的区别，因为单个插槽和具名插槽不绑定数据，所以父组件是提供的模板要既包括样式由包括内容的，上面的例子中，你看到的文字，“菜单1”，“菜单2”都是父组件自己提供的内容；而作用域插槽，父组件只需要提供一套样式（在确实用作用域插槽绑定的数据的前提下）。

下面的例子，你就能看到，父组件提供了三种样式(分别是flex、ul、直接显示)，都没有提供数据，数据使用的都是子组件插槽自己绑定的那个人名数组。

父组件：

```HTML
<template>
  <div class="child">

    <h3>这里是子组件</h3>

    <slot  :data="data"></slot>
  </div>
</template>

<script>
  export default {
      data: function(){
        return {
          data: ['zhangsan','lisi','wanwu','zhaoliu','tianqi','xiaoba']
        }
      },
    computed: {
    },
    methods:{
    },
    components: {
    }
  }
</script>

<style scoped>
  .child{
    background-color: #00bbff;
    width: 100%;
    padding: 10px;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
</style>

```

子组件：

```HTML
<template>
  <div class="child">

    <h3>这里是子组件</h3>

    <slot  :data="data"></slot>
  </div>
</template>

<script>
  export default {
      data: function(){
        return {
          data: ['zhangsan','lisi','wanwu','zhaoliu','tianqi','xiaoba']
        }
      },
    computed: {
    },
    methods:{
    },
    components: {
    }
  }
</script>

<style scoped>
  .child{
    background-color: #00bbff;
    width: 100%;
    padding: 10px;
    -webkit-box-sizing: border-box;
    -moz-box-sizing: border-box;
    box-sizing: border-box;
  }
</style>
```

结果如图所示：

![img](https://segmentfault.com/img/remote/1460000012996224?w=703&h=651)



以上三个demo就放在https://github.com/cunzaizhuyi/vue-slot-demo了，有需要的可以去取。使用非常方便，是基于vue-cli搭建工程。

### 笔记不易：

离线笔记均为markdown格式，图片也是云图，10多篇笔记20W字，压缩包仅500k，推荐使用typora阅读。也可以自己导入有道云笔记等软件中

阿里云图床现在**每周得几十元充值**，都要自己往里搭了，麻烦不要散播与转发

![](https://i0.hdslb.com/bfs/album/ff3fb7e24f05c6a850ede4b1f3acc54312c3b0c6.png)

打赏后请主动发支付信息到邮箱  553736044@qq.com  ，上班期间很容易忽略收账信息，邮箱回邮基本秒回（请备注付费的内容）

禁止转载发布，禁止散播，若发现大量散播，将对本系统文章图床进行重置处理。

技术人就该干点技术人该干的事



如果帮到了你，留下赞吧，谢谢支持