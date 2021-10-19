# jelly: jelly fish ssh editor
## auth: liikii


#### init go files
```shell
mkdir jelly
cd jelly
go mod init example.com/m
```



#### init web files
```shell
npx degit sveltejs/template jelly-app
cd jelly-app
npm install
npm run dev
npm run build
```



#### svelte import css
```html
<style>
	@import url("https://stackpath.bootstrapcdn.com/bootstrap/4.5.0/css/bootstrap.min.css");
</style>
```