<template>
	<div class="login-view">
		<div class="container">
			<div :class="shortContainerClasses">
				<div>
					<h2>Log In</h2>
					<form @submit.prevent="submitCreds()">
						<input
							type="email"
							placeholder="Email address"
							v-model="email">
						<input
							type="password"
							placeholder="Password"
							v-model="password">
						<div class="submit-area">
							<div class="status">{{ status || '' }}</div>
							<button type="submit">Log in</button>
						</div>
					</form>
				</div>

				<div class="note">
					<p><strong>Note:</strong></p>
					<p>This is a concept site; there is no way for a new user to register an account and log in.</p>
					<p>This page is simply available for administrative tasks, such as adding or removing content.</p>
					<p>If you'd like to contact the developer of this site, please visit <a href="https://jamiesyme.com">jamiesyme.com</a>.</p>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	import Auth from '@/auth';

	export default {
		data () {
			return {
				email: '',
				password: '',
				status: null,
			};
		},

		mounted () {
			document.title = 'Login | Videoso';
		},

		methods: {
			goToAdmin () {
				this.$router.push('admin');
			},

			submitCreds () {
				if (Auth.login(this.email, this.password)) {
					this.goToAdmin();
				} else {
					// Add a slight delay so that on subsequent attempts there is
					// still an indication that the credentials are being tried.
					this.status = null;
					setTimeout(() => {
						this.status = 'Invalid credentials';
					}, 250);
				}
			},
		},

		computed: {
			shortContainerClasses () {
				const bp = this.$breakpoint.name;
				const half = bp !== 'small';
				return {
					'short-container': true,
					'short-container-half': half,
				};
			},
		},
	}
</script>

<style lang="scss" scoped>
	.short-container {
		margin: 4rem auto 6rem;

		&.short-container-half {
			width: 50%;
		}
	}

	h2 {
		margin: 0 0 2rem;
	}

	.submit-area {
		display: flex;

		.status {
			flex: 1;
			line-height: 3.8rem;
			font-weight: bold;
			font-size: 1.4rem;
			color: #ff0033;
		}

		button {
			display: block;
			width: 10rem;
			margin: 0 0 0 auto;
		}
	}

	.note {
		//background-color: #ece6f2;
		//background-color: #DEF3FA;
		background-color: #E6F6ED;
		//border: 0.2rem solid #9b4dca;
		//border: 0.2rem solid #40bae3;
		border: 0.2rem solid #60C78F;
		border-radius: 0.4rem;
		padding: 1.6rem;

		p:last-of-type {
			margin-bottom: 1rem;
		}

		a {
			color: #60C78F;
			font-weight: bold;
		}
	}
</style>
