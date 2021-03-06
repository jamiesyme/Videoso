<template>
	<div class="admin-view container">
		<div class="sidebar">
			<nav>
				<router-link
					class="button button-clear"
					to="/admin/categories">
					Categories
				</router-link>
				<router-link
					class="button button-clear"
					to="/admin/users">
					Users
				</router-link>
				<router-link
					class="button button-clear"
					to="/admin/videos">
					Videos
				</router-link>
			</nav>
			<hr>
			<button
				class="button button-clear"
				:disabled="saving || loading"
				@click="saveContent()">
				Save
			</button>
		</div>

		<div class="sub-view">
			<section
				class="categories-section"
				v-if="section === 'categories'">
				<h2>Categories</h2>
				<table>
					<thead>
						<tr>
							<th class="title-col">Title</th>
							<th class="delete-col">Delete</th>
						</tr>
					</thead>
					<tbody>
						<tr
							class="category"
							v-for="category in categories">
							<td>
								<LabelInput v-model="category.title" />
							</td>
							<td>
								<button
									class="button-clear"
									@click="removeCategory(category.id)">
									<IonClose />
								</button>
							</td>
						</tr>
						<tr class="new-category">
							<td>
								<button
									class="button-clear"
									@click="addCategory()">
									<IonAdd />Add Category
								</button>
							</td>
							<td></td>
						</tr>
					</tbody>
				</table>
			</section>

			<section
				class="users-section"
				v-if="section === 'users'">
				<h2>Users</h2>
				<table>
					<thead>
						<tr>
							<th class="name-col">Name</th>
							<th class="banner-col">Banner URL</th>
							<th class="delete-col">Delete</th>
						</tr>
					</thead>
					<tbody>
						<tr
							class="user"
							v-for="user in users">
							<td>
								<LabelInput v-model="user.name" />
							</td>
							<td class="banner-td">
								<LabelInput v-model="user.bannerUrl">
									<template v-slot:empty>
										<span class="null-banner">none</span>
									</template>
								</LabelInput>
							</td>
							<td>
								<button
									class="button-clear"
									@click="removeUser(user.id)">
									<IonClose />
								</button>
							</td>
						</tr>
						<tr class="new-user">
							<td>
								<button
									class="button-clear"
									@click="addUser()">
									<IonAdd />Add User
								</button>
							</td>
							<td></td>
						</tr>
					</tbody>
				</table>
			</section>

			<section
				class="videos-section"
				v-if="section === 'videos'">
				<h2>Videos</h2>
				<table>
					<thead>
						<tr>
							<th class="video-col">Video</th>
							<th class="delete-col">Delete</th>
						</tr>
					</thead>
					<tbody>
						<tr
							class="video"
							v-for="video in videos">
							<td>
								<table>
									<tr>
										<td>Title</td>
										<td>
											<LabelInput v-model="video.title" />
										</td>
									</tr>
									<tr>
										<td>Description</td>
										<td>
											<LabelInput
												type="textarea"
												v-model="video.description" />
										</td>
									</tr>
									<tr>
										<td>Tags</td>
										<td>
											<LabelInput v-model="video.tags" />
										</td>
									</tr>
									<tr>
										<td>Published at</td>
										<td>
											<LabelInput
												type="datetime-local"
												v-model="video.publishedAt">
												<template v-slot:label>
													<div class="label">
														{{ prettyDateTime(video.publishedAt) }}
													</div>
												</template>
											</LabelInput>
										</td>
									</tr>
									<tr>
										<td>View count</td>
										<td>
											<LabelInput
												type="number"
												v-model="video.viewCount" />
										</td>
									</tr>
									<tr>
										<td>Category</td>
										<td>
											<LabelInput
												list="category-datalist"
												v-model="video.category" />
										</td>
									</tr>
									<tr>
										<td>Author</td>
										<td>
											<LabelInput
												list="user-datalist"
												v-model="video.author" />
										</td>
									</tr>
									<tr>
										<td>Duration</td>
										<td>
											<LabelInput
												type="text"
												v-model="video.duration" />
										</td>
									</tr>
									<tr>
										<td>Video URL</td>
										<td>
											<LabelInput
												type="text"
												v-model="video.videoUrl" />
										</td>
									</tr>
									<tr>
										<td>Thumbnail URL</td>
										<td>
											<LabelInput
												type="text"
												v-model="video.thumbnailUrl" />
										</td>
									</tr>
									<tr>
										<td>Responsive Thumbnails</td>
										<td>
											<input
												type="checkbox"
												v-model="video.responsiveThumbnails">
										</td>
									</tr>
								</table>
							</td>
							<td class="delete-td">
								<button
									class="button-clear"
									@click="removeVideo(video.id)">
									<IonClose />
								</button>
							</td>
						</tr>
						<tr class="new-video">
							<td>
								<button
									class="button-clear"
									@click="addVideo()">
									<IonAdd />Add Video
								</button>
							</td>
							<td></td>
						</tr>
					</tbody>
				</table>

				<datalist id="category-datalist">
					<option
						:value="category.title"
						v-for="category in categories" />
				</datalist>
				<datalist id="user-datalist">
					<option
						:value="user.name"
						v-for="user in users" />
				</datalist>
			</section>
		</div>
	</div>
</template>

<script>
	import IonAdd from 'vue-ionicons/dist/md-add';
	import IonClose from 'vue-ionicons/dist/md-close';
	import LabelInput from '@/components/LabelInput';
	import Content from '@/content';

	export default {
		props: {
			section: {
				type: String,
				default: 'categories',
			},
		},

		components: {
			IonAdd,
			IonClose,
			LabelInput,
		},

		data () {
			return {
				loading: false,
				saving: false,
				categories: [
					/**
					 * @typedef {object} Category
					 * @property {number} id
					 * @property {string} title
					 */
				],
				users: [
					/**
					 * @typedef {object} User
					 * @property {number} id
					 * @property {string} name
					 * @property {string} [bannerUrl]
					 */
				],
				videos: [
					/**
					 * @typedef {object} Video
					 * @property {number}  id
					 * @property {string}  title
					 * @property {string}  description
					 * @property {string}  tags - space-separated
					 * @property {string}  publishedAt - YYYY-MM-DDThh:mm
					 * @property {number}  viewCount
					 * @property {string}  category - category.title
					 * @property {string}  author - user.name
					 * @property {string}  duration - mm:ss
					 * @property {string}  videoUrl
					 * @property {string}  thumbnailUrl
					 * @property {boolean} responsiveThumbnails
					 */
				],
			};
		},

		async mounted () {
			document.title = 'Admin | Videoso';
			await this.loadContent();
		},

		methods: {
			async loadContent () {
				this.loading = true;
				await Content.load({ fresh: true });
				this.categories = Content.categories.slice();
				this.users = Content.users.slice();
				this.videos = Content.videos.map(v => {
					// When we load videos, we have to translate a couple fields.
					// Sometimes this involves denormalizing info (as is the case
					// with the `category` field), but in return, we can re-use
					// existing input elements that work with simple strings,
					// etc. rather than building custom input elements for each
					// type of input.
					const pub = new Date(Date.parse(v.publishedAt));
					const cat = this.categories.find(c => c.id === v.category);
					const usr = this.users.find(u => u.id === v.author);
					const dur = this.durationToStr(v.duration);
					return {
						id:           v.id,
						title:        v.title,
						description:  v.description,
						tags:         v.tags.join(' '),
						publishedAt:  pub,
						viewCount:    v.viewCount,
						category:     cat.title,
						author:       usr.name,
						duration:     dur,
						videoUrl:     v.videoUrl,
						thumbnailUrl: v.thumbnailUrl,
						responsiveThumbnails: v.responsiveThumbnails,
					};
				});
				this.loading = false;
			},

			async saveContent () {
				this.saving = true;
				Content.categories = this.categories.slice();
				Content.users = this.users.slice();
				Content.videos = this.videos.map(v => {
					// While editing videos we use an intermediate format (see
					// justification above), so before we save the videos we have
					// to translate all of the fields back to their normalized
					// values.
					const pub = v.publishedAt.toISOString();
					const cat = this.findCategoryByTitle(v.category);
					const usr = this.findUserByName(v.author);
					const dur = this.strToDuration(v.duration);
					return {
						id:           v.id,
						title:        v.title,
						description:  v.description,
						tags:         v.tags.split(' '),
						publishedAt:  pub,
						viewCount:    v.viewCount,
						category:     cat.id,
						author:       usr.id,
						duration:     dur,
						videoUrl:     v.videoUrl,
						thumbnailUrl: v.thumbnailUrl,
						responsiveThumbnails: v.responsiveThumbnails,
					};
				});
				await Content.save();
				this.saving = false;
			},

			addCategory () {
				const nextId = this.categories.reduce((accum, next) => {
					return Math.max(next.id + 1, accum);
				}, 1);
				this.categories.push({
					id: nextId,
					title: 'New Category',
				});
			},

			removeCategory (categoryId) {
				const index = this.categories.findIndex(c => {
					return c.id === categoryId;
				});
				this.categories.splice(index, 1);
			},

			findCategoryByTitle (title) {
				const s1 = title.toLowerCase().trim();
				return this.categories.find(category => {
					const s2 = category.title.toLowerCase().trim();
					return s1 === s2;
				});
			},

			addUser () {
				const nextId = this.users.reduce((accum, next) => {
					return Math.max(next.id + 1, accum);
				}, 1);
				this.users.push({
					id: nextId,
					name: 'New User',
					bannerUrl: null,
				});
			},

			removeUser (userId) {
				const index = this.users.findIndex(c => {
					return c.id === userId;
				});
				this.users.splice(index, 1);
			},

			findUserByName (name) {
				const s1 = name.toLowerCase().trim();
				return this.users.find(user => {
					const s2 = user.name.toLowerCase().trim();
					return s1 === s2;
				});
			},

			addVideo () {
				const nextId = this.videos.reduce((accum, next) => {
					return Math.max(next.id + 1, accum);
				}, 1);
				this.videos.push({
					id: nextId,
					title: 'New Video',
					description: '',
					tags: '',
					publishedAt: new Date,
					viewCount: 0,
					category: null,
					author: null,
					duration: '00:00',
					videoUrl: null,
					thumbnailUrl: null,
					responsiveThumbnails: false,
				});
			},

			removeVideo (videoId) {
				const index = this.videos.findIndex(c => {
					return c.id === videoId;
				});
				this.videos.splice(index, 1);
			},

			prettyDateTime (date) {
				return date.toLocaleString('en-US', {
					month: 'short',
					day: 'numeric',
					year: 'numeric',
					hour: 'numeric',
					minute: '2-digit',
					hour12: true,
				});
			},

			durationToStr (seconds) {
				function pad2 (num) {
					return num.toString().padStart(2, '0');
				}
				const secs = Math.floor(seconds % 60);
				const mins = Math.floor(seconds / 60);
				const hrs = Math.floor(seconds / 60 / 60);
				if (hrs > 0) {
					return `${pad2(hrs)}:${pad2(mins)}:${pad2(secs)}`;
				} else {
					return `${pad2(mins)}:${pad2(secs)}`;
				}
			},

			strToDuration (str) {
				const parts = str.split(':');
				let hrs = 0;
				let mins = 0;
				let secs = 0;
				switch (parts.length) {
					case 1:
						secs = parseInt(parts[0]);
						break;
					case 2:
						mins = parseInt(parts[0]);
						secs = parseInt(parts[1]);
						break;
					case 3:
						hrs = parseInt(parts[0]);
						mins = parseInt(parts[1]);
						secs = parseInt(parts[2]);
						break;
					default:
						throw new Error();
				}
				return (hrs * 60 + mins) * 60 + secs;
			},
		},

		computed: {
			// Exists solely so that we can easily watch for changes to category
			// titles.
			categoriesByTitle () {
				const obj = {};
				for (const cat of this.categories) {
					obj[cat.title] = cat.id;
				}
				return obj;
			},

			// Exists solely so that we can easily watch for changes to user
			// names.
			usersByName () {
				const obj = {};
				for (const user of this.users) {
					obj[user.name] = user.id;
				}
				return obj;
			},
		},

		watch: {
			categoriesByTitle: function (newValues, oldValues) {
				// Videos track their category by title, not id, so we need to
				// keep those values up to date when a category's name changes.
				for (const video of this.videos) {
					const catId = oldValues[video.category];
					const cat = this.categories.find(c => c.id === catId);
					if (cat) {
						video.category = cat.title;
					}
				}
			},

			usersByName: function (newValues, oldValues) {
				// Videos track their author by name, not id, so we need to keep
				// those values up to date when a user's name changes.
				for (const video of this.videos) {
					const userId = oldValues[video.author];
					const user = this.users.find(u => u.id === userId);
					if (user) {
						video.author = user.name;
					}
				}
			},
		},
	}
</script>

<style lang="scss" scoped>
	.admin-view {
		display: flex;
		margin: 4rem auto 6rem;
		padding: 0;
	}

	.sidebar {
		width: 25%;
		border-right: 0.1rem solid #eee;
		padding-right: 3rem;
		margin-right: 3rem;

		.button {
			display: block;
			width: 100%;
			text-align: left;
			padding: 0 1rem;
		}
	}

	.sub-view {
		flex: 1;
		margin: 0 3rem 0 0;

		.categories-section, .users-section, .videos-section {
			table {
				button {
					padding: 0;
					margin: 0;
				}
				.ion {
					font-size: 2.2rem;
				}
				.category, .user, .video {
					button {
						width: 100%;
					}
					.ion {
						display: block;
					}
				}
				.new-category, .new-user, .new-video {
					.ion {
						position: relative;
						top: 0.6rem;
						margin: 0 0.6rem 0 0;
					}
				}
			}
		}

		.categories-section {
			.title-col {
				width: 100%;
			}
		}

		.users-section {
			.name-col {
				width: 50%;
			}
			.banner-col {
				width: 50%;
			}
			.banner-td {
				max-width: 50rem;
			}
			.null-banner {
				font-style: italic;
			}
		}

		.videos-section {
			.video-col {
				width: 100%;
			}
			.delete-td {
				vertical-align: top;
			}
			.video table {
				padding-left: 2.5rem;
				margin-bottom: 0;
				table-layout: fixed;
				td {
					border: none;
				}
				td:first-of-type {
					width: 20rem;
				}
				input[type="checkbox"] {
					display: block;
					width: 2.4rem;
					height: 2.4rem;
					margin: 0.2rem 0;
				}
			}
		}
	}
</style>
