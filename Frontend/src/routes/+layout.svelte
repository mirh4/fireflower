<script lang="ts">
    let { children } = $props();
</script>

<nav>
    <a href="/">Home</a>
    <a href="/about">About</a>
    <a href="/settings">Settings</a>
    <a href="/blog">Blog</a>
</nav>

{@render children()}