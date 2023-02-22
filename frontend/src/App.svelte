<script lang="ts">
    import {AppShell, AppBar} from '@skeletonlabs/skeleton'
    import {OpenDirDialog, ReadDir, ReadFile} from '../wailsjs/go/main/App.js'
    import {each} from "svelte/internal";

    let eliteDir: string = ""
    let files: string[] = []
    $: eliteDir && readDir(eliteDir)

    function chooseDir(): void {
        OpenDirDialog().then(dir => eliteDir = dir)
    }

    function readDir(dir): void {
        console.log(dir)
        if (dir) {
            ReadDir(dir).then(journal_files => {
                files = journal_files
                console.log(files)
            })
        }
    }
    function readFile(file) {
        console.log(file)
        if (file) {
            ReadFile(file).then(events => {
                each(events, event => {
                    console.log(JSON.parse(event))
                })
            })
        }
    }
</script>
<AppShell>
    <svelte:fragment slot="header"><AppBar>Go Loopy Science!</AppBar></svelte:fragment>
    <svelte:fragment slot="pageHeader"><h1 class="text-center">Step 1</h1></svelte:fragment>
    <main>
        <div class="flex justify-evenly items-center">
            <div class="">{eliteDir ? eliteDir : 'Please choose your Elite Dangerous journal directory.'}</div>
            <button class="btn variant-ghost-primary" on:click={chooseDir}>Browse</button>
        </div>

        {#if files.length > 0}
            <ul>
                {#each files as file}
                    <li on:click={readFile(file)}>{file}</li>
                {/each}
            </ul>
        {/if}
    </main>
    <svelte:fragment slot="pageFooter">Page Footer</svelte:fragment>
    <svelte:fragment slot="footer">Footer</svelte:fragment>
</AppShell>


<style>


</style>
