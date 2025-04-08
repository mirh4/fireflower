// @ts-nocheck
import { error } from '@sveltejs/kit';
import type { PageServerLoad } from './$types.d.ts';

export const load = async ({ params }: Parameters<PageServerLoad>[0]) => {
	const post = await getPostFromDatabase(params.slug);

	if (post) {
		return post;
	}

	error(404, 'Not found');
};