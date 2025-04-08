// @ts-nocheck
import { error } from '@sveltejs/kit';
import type { PageLoad } from './$types.d.ts';

export const load = ({ params }: Parameters<PageLoad>[0]) => {
	if (params.slug === 'hello-world') {
		return {
			title: 'Hello world!',
			content: 'Welcome to our blog. Lorem ipsum dolor sit amet...'
		};
	}

	error(404, 'Not found');
};