<script>
</script>

<header>
	<div class="corner">
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L0,3 C0.5,3 0.5,3 1,2 L2,0 Z" />
		</svg>
	</div>
	<nav>
		<img src="favicon.png" alt="favicon" width="80" height="80" />
		<h1 class="display-1">Flag Fairway</h1>
	</nav>
	<div class="corner">
		<svg viewBox="0 0 2 3" aria-hidden="true">
			<path d="M0,0 L1,2 C1.5,3 1.5,3 2,3 L2,0 Z" />
		</svg>
	</div>
	
</header>

<style lang="sass">
$background: rgba(255, 255, 255, 0.7)

$color-theme-1: #your-theme-color
$color-text: #your-text-color

header
  display: flex
  justify-content: space-between

  .corner
    width: 3em
    height: 3em

    img
      width: 2em
      height: 2em
      object-fit: contain

nav
  display: flex
  justify-content: center

svg
  width: 2em
  height: 3em
  display: block

  path
    fill: $background
</style>
