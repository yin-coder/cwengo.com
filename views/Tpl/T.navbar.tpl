<div class="navbar-header">
			<button type="button" class="navbar-toggle" data-toggle="collapse" 
		         data-target="#example-navbar-collapse">
		         <span class="sr-only">切换导航</span>
		         <span class="icon-bar"></span>
		         <span class="icon-bar"></span>
		         <span class="icon-bar"></span>
		     </button>
		     <a class="navbar-brand" href="/" ><big>沉稳，不乏可爱</big></a>	
		</div>
		<div>
			<p class="navbar-text">一个IT技术宅的私家宅地</p>
		</div>
		<div class="pull-right">
			<ul class="nav navbar-nav">
				<li {{if .IsHome}} class="active" {{end}} ><a href="/">主页</a></li>
				<li ><a href="/">文章归档</a></li>
				<li ><a href="/">博主</a></li>
				<li></li>
			</ul>
		</div>