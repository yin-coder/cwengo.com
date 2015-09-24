<link rel="stylesheet"  type="text/css" href="/static/css/admin.css">
<script type="text/javascript" src="/static/js/category.js"></script>
<div class="row-fluid">
	<div class="btn-toolbar">
		<button type="button" class="btn btn-primary" data-toggle="modal" data-target="#addCategory"><span class="glyphicon glyphicon-plus" aria-hidden="true"  ></span>添加分类</button>
	</div>
	<div class="well">
		<table class="table table-hover">
		 	<thead>
		 		<th>分类ID</th>
		 		<th>分类名称</th>
		 		<th>文章数</th>
		 		<th>操作</th>
		 	</thead>
		 	<tbody>
		 		{{range .Categories}}
		 		<tr>
		 			<th>{{.Id}}</th>
		 			<th>{{.Title}}</th>
		 			<th>{{.TopicCount}}</th>
		 			<td>
		 				<a href=""><span class="glyphicon glyphicon-pencil" aria-hidden="true"></span></a>&nbsp;&nbsp;
		 				<a href=""><span class="glyphicon glyphicon-trash" aria-hidden="true"></span></a>
		 			</td>
		 		</tr>
		 		{{end}}
		 	</tbody>
		</table>
	</div>
</div>

<div class="modal fade" id="addCategory" tabindex="-1" role="dialog" aria-labelledby="myModalLabel">
  <div class="modal-dialog" role="document">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
        <h4 class="modal-title" id="myModalLabel">添加分类</h4>
      </div>
      <div class="modal-body">
       		<form  method="post">
       		  <div class="form-group">
       		    <label>新添分类名</label>
       		    <input type="text" class="form-control" name="NewCategory" id="NewCategory" placeholder="Email">
       		  </div>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">关闭</button>
        <button type="submit" class="btn btn-primary" onclick="return checkInput();" >添加</button>
        <script>
        function checkInput() {
	        var NewCategory = document.getElementById("NewCategory");
	        if(NewCategory.value.length == 0) {
	            alert("分类名称不能为空");
	            return false;
	        }
      	}
        </script>
        </form>
      </div>
    </div>
  </div>
</div>