<script>
	// export let name;
	import { onMount } from 'svelte';
	import LeftArrow from './LeftArrow.svelte';
	import ScrollText from './ScrollText.svelte';
	import LoginPop from './LoginPop.svelte';
	// let login_model = "display: block;";
	// login_model = "display: none;";
	// let login_model;
	// let modal_show = "block";

	// let loggedIn = false;
	// if (!loggedIn) {
	// 	login_model.modal('show');
	// }
	// onMount( () => {
	// 	console.log(login_model);
		// if (!loggedIn) {
		// 	login_model.show();
		// }
	// });
	let show_login_form = "block";
	function logout() {
		// body...
		show_login_form = "block";
	}

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
	        	<div class="col-2 py-1 px-0 bg-primary">
	        		<LeftArrow/>
	        	</div>
	        		<ScrollText dev_css="col-9 p-0" roll_text="ddddddddddddd/home"/>
	      	</div>
	      </div>
	      
	      <div class="col-sm-9 col-md-6 col-lg-8 col-xl-10 bg-warning">
	        菜鸟教程
	      </div>
	    </div>
	  </div>

	  <button on:click={logout}>登出</button>
	  <LoginPop bind:display_v={show_login_form}/>
	  <!-- <div class="modal show" style="{login_model}" tabindex="1" role="dialog">
		  <div class="modal-dialog" role="document">
		    <div class="modal-content">
		      <div class="modal-header">
		        <h5 class="modal-title">Modal title</h5>
		        <button type="button" class="close" data-dismiss="modal" aria-label="Close">
		          <span aria-hidden="true">&times;</span>
		        </button>
		      </div>
		      <div class="modal-body">
		        <p>Modal body text goes here.</p>
		      </div>
		      <div class="modal-footer">
		        <button type="button" class="btn btn-primary">Save changes</button>
		        <button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
		      </div>
		    </div>
		  </div>
		</div> -->

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