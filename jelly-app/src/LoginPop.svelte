<script>
	export let display_v = "block";
	function handleClick() {
		console.log('no more alerts');
		display_v= "none";
	}
</script>

<div class="fixed-top bg-primary h-100 w-100" style="display: {display_v};">
	<div class="d-flex justify-content-center">
	  <div class="p-2">Flex item 1</div>
	  <div class="p-2">Flex item 2</div>
	  <div class="p-2">Flex item 3</div>
	  <button type="button" on:click={handleClick}>
	    登陆
	  </button>
	</div>
</div>
