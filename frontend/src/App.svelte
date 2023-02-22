<script lang="ts">

  import {OpenDirDialog} from '../wailsjs/go/main/App.js'

  let eliteDir: string = "Please choose your Elite Dangerous journal directory."



  function chooseDir(): void {
    OpenDirDialog().then(dir => eliteDir = dir)
  }
</script>

<main>

  <div class="result" id="result">{eliteDir}</div>
  <div class="input-box" id="input">
    <button class="btn" on:click={chooseDir}>Browse</button>
  </div>
</main>

<style>

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

</style>
