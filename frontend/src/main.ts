// Your selected Skeleton theme:
import '@skeletonlabs/skeleton/themes/theme-crimson.css';

// This contains the bulk of Skeletons required styles:
import '@skeletonlabs/skeleton/styles/all.css';
import './app.css'
import App from './App.svelte'

const app = new App({
  target: document.getElementById('app')
})

export default app
