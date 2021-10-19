<script>
	// export let name;
	import LeftArrow from './LeftArrow.svelte';
</script>

<main>
<!-- 	<h1 class="display-1">Hello {name}!</h1>
	<p class="text-success">Visit the https://svelte.dev/tutorial to learn how to build Svelte apps.</p> -->
	<div class="container-fluid">
	  <div class="container-fluid">
	    <div class="row">
	      <div class="col-sm-3 col-md-6 col-lg-4 col-xl-2">
	      	
	        <LeftArrow/>
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