<script>
    import Header from './Header.svelte'
	import './styles.sass'
</script>

<div class="app">
	<Header />
	<main>
		<slot />
	</main>
	<footer>
		<p>Copyright Luke Liukonen 2024</p>
	</footer>
</div>

<style lang="sass">

.app
  display: flex
  flex-direction: column
  min-height: 100vh

main
  flex: 1
  display: flex
  flex-direction: column
  padding: 1rem
  width: 100%
  max-width: 64rem
  margin: 0 auto
  box-sizing: border-box

footer
  display: flex
  flex-direction: column
  justify-content: center
  align-items: center
  padding: 12px

@media (min-width: 480px)
  footer
    padding: 12px 0
</style>