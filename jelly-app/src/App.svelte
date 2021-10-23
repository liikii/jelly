<script>
	// export let name;
	import LeftArrow from './LeftArrow.svelte';
	import ScrollText from './ScrollText.svelte';
</script>

<main>
<!-- 	<h1 class="display-1">Hello {name}!</h1>
	<p class="text-success">Visit the https://svelte.dev/tutorial to learn how to build Svelte apps.</p> -->
	<div class="container-fluid">
	  <div class="container-fluid">
	    <div class="row">
	      <div class="col-sm-3 col-md-6 col-lg-4 col-xl-2 p-0">
	      	<div class="row">
	        	192.168.18.9
	      	</div>
	      	<!-- <LeftArrow/> -->
	      	<div class="row px-1 bg-primary" style="height: 50px;">
	        	<div class="col-2 p-0 bg-primary">
	        		<LeftArrow/>
	        	</div>
	        		<ScrollText dev_css="col-8 p-0" roll_text="ddddddddddddd/home"/>
	      	</div>
	      </div>
	      
	      <div class="col-sm-9 col-md-6 col-lg-8 col-xl-10 bg-warning">
	        菜鸟教程
	      </div>
	    </div>
	  </div>
	</div>
</main>

<style>
	/*main {
		text-align: center;
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}*/

/*	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}
*/
	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>