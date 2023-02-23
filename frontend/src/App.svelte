<script lang="ts">
    import {AppShell, AppBar} from '@skeletonlabs/skeleton'
    import {OpenDirDialog} from '../wailsjs/go/main/App.js'

    let journalDir: string = ""

    function chooseDir(): void {
        OpenDirDialog().then(dir => journalDir = dir)
    }


</script>
<AppShell>
    <svelte:fragment slot="header">
        <AppBar>Go Loopy Science!</AppBar>
    </svelte:fragment>
    <svelte:fragment slot="pageHeader"><h1 class="text-center">Step 1 <small class="text-base text-primary-500">Auto detecting Journal Directory</small></h1></svelte:fragment>
    <svelte:fragment slot="pageContent">
        <div class="flex flex-col items-center justify-center h-full">
            {#if journalDir}
                <h2 class="text-center">Journal Directory: {journalDir}</h2>
            {:else}
                <h2 class="text-center">Journal Directory not found</h2>
            {/if}
        </div>
    </svelte:fragment>
    <svelte:fragment slot="pageFooter">Page Footer</svelte:fragment>
    <svelte:fragment slot="footer">Footer</svelte:fragment>
</AppShell>


<style>


</style>
