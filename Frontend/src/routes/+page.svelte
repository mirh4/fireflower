<script>
    import { onMount } from 'svelte';

    let notes = [];
    let title = '';
    let content = '';

    // Fetch notes on load
    onMount(async () => {
        const res = await fetch('http://localhost:8080/notes');
        if (res.ok) {
            notes = await res.json();
        }
    });

    // Create a new note
    async function addNote() {
        const res = await fetch('http://localhost:8080/notes', {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ title, content })
        });
        if (res.ok) {
            const newNote = await res.json();
            notes = [...notes, newNote];
            title = '';
            content = '';
        }
    }

    // Delete a note
    async function deleteNote(id) {
        const res = await fetch(`http://localhost:8080/notes/${id}`, {
            method: 'DELETE'
        });
        if (res.ok) {
            notes = notes.filter(note => note.id !== id);
        }
    }
</script>

<main>
    <h1>Notes App</h1>

    <div>
        <input bind:value={title} placeholder="Note Title" />
        <textarea bind:value={content} placeholder="Note Content"></textarea>
        <button on:click={addNote}>Add Note</button>
    </div>

    <h2>Your Notes</h2>
    {#each notes as note}
        <div class="note">
            <h3>{note.title}</h3>
            <p>{note.content}</p>
            <small>{note.created_at}</small>
            <button on:click={() => deleteNote(note.id)}>Delete</button>
        </div>
    {/each}
</main>

<style>
    .note {
        border: 1px solid #ccc;
        padding: 10px;
        margin: 10px 0;
    }
    button {
        margin-top: 5px;
        background-color: #ff4d4d;
        color: white;
        border: none;
        padding: 5px 10px;
        cursor: pointer;
    }
    button:hover {
        background-color: #cc0000;
    }
</style>