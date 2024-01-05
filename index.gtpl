<!DOCTYPE html>
<html>
<head>
    <title>File Upload</title>
    <style>
        /* 添加一些基本的样式和你的模态 */
        .modal {
            display: none;
            position: fixed;
            z-index: 1;
            left: 0;
            top: 0;
            width: 100%;
            height: 100%;
            overflow: auto;
            background-color: rgb(0,0,0);
            background-color: rgba(0,0,0,0.4);
        }
        .modal-content {
            background-color: #fefefe;
            margin: 15% auto;
            padding: 20px;
            border: 1px solid #888;
            width: 30%;
        }
        .close {
            color: #aaa;
            float: right;
            font-size: 28px;
            font-weight: bold;
        }
        .close:hover,
        .close:focus {
            color: black;
            text-decoration: none;
            cursor: pointer;
        }
        .submit-button {
            color: white;
            background-color: #007aff;
            border: none;
            padding: 10px 20px;
            text-align: center;
            text-decoration: none;
            display: inline-block;
            font-size: 16px;
            margin: 10px 2px;
            cursor: pointer;
        }
    </style>
</head>
<body>
    <h1>Current Directory: {{.Dir}}</h1>
    <h2><a href="{{.ParentDir}}">&#8678; Back</a></h2>
    <ul>
    {{range .Files}}
        <li><a href="/{{.Path}}">{{.Name}}</a></li>
    {{end}}
    </ul>

    <!-- 将上传表单放在模态框中 -->
    <div id="myModal" class="modal">
        <div class="modal-content">
            <span class="close">&times;</span>
            <form action="/{{.Dir}}" method="post" enctype="multipart/form-data">
                Select file to upload:
                <input type="file" name="file">
                <input class="submit-button" type="submit" value="Upload">
            </form>
        </div>
    </div>

    <button id="myBtn" class="submit-button">Upload File</button>

    <script>
        // 获取模态框元素
        var modal = document.getElementById("myModal");

        // 获取开启模态框的按钮元素
        var btn = document.getElementById("myBtn");

        // 获取模态框上的 <span> 元素, 用于关闭模态框
        var span = document.getElementsByClassName("close")[0];

        // 在用户点击按钮时，打开模态框
        btn.onclick = function() {
            modal.style.display = "block";
        }

        // 在用户点击 <span> (x), 关闭模态框
        span.onclick = function() {
            modal.style.display = "none";
        }

        // 在用户点击其他地方，也关闭模态框
        window.onclick = function(event) {
            if (event.target === modal) {
                modal.style.display = "none";
            }
        }
    </script>
</body>
</html>