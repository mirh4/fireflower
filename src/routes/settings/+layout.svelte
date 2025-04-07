<script lang="ts">
	import type { LayoutProps } from './$types.d.ts';

	let { data, children }: LayoutProps = $props();
</script>

<h1>Settings</h1>

<div class="submenu">
	{#each data.sections as section}
		<a href="/settings/{section.slug}">{section.title}</a>
	{/each}
</div>

{@render children()}
