<script lang="ts">
	import type { PageProps } from './$types.d.ts';

	let { data }: PageProps = $props();
</script>

<h1>{data.title}</h1>
<div>{@html data.content}</div>