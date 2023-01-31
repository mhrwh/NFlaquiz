<template>
	<div class="modal fade" id="helpModal" tabindex="-1" aria-labelledby="helpModalTitle" aria-hidden="false">
		<div class="modal-dialog modal-dialog-centered">
			<div class="modal-content help-modal">
				<div class="modal-header p-0 border-bottom-0">
					<h2 class="modal-title" id="helpModalTitle">ヘルプ</h2>
				</div>
				<div class="modal-body p-0">
					<ul class="nav nav-tabs">
						<li class="nav-item">
							<a class="nav-link active" data-toggle="tab" href="#tab1" @click="selectedTab = 'tab1'">
								ボタンの説明
							</a>
						</li>
						<li class="nav-item">
							<a class="nav-link" data-toggle="tab" href="#tab2" @click="selectedTab = 'tab2'">
								ログイン機能
							</a>
						</li>
					</ul>
					<div class="tab-content">
						<div class="tab-pane fade show active" id="tab1" v-if="selectedTab === 'tab1'">
							<i class="bi bi-person-fill" />
							ログインボタン<br>
							<i class="bi bi-layout-sidebar-inset" />
							マップ切り替えボタン（正答率モード）<br>
							<i class="bi bi-layout-sidebar-inset-reverse" />
							マップ切り替えボタン（ブックマークモード）<br>
							<i class="bi bi-bookmark-heart" />
							ブックマーク編集ボタン<br>
							<i class="bi bi-flag-fill" />
							クイズボタン
						</div>
						<div class="tab-pane fade" id="tab2" v-if="selectedTab === 'tab2'">
							ログインすることでブックマーク機能が使えたり、<br>
							正答率・ブックマークした国をマップに反映することができます。
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
export default {
	data() {
		return {
			showMore: false,
			selectedTab: 'tab1'
		}
	}
}
</script>

<style>
.modal-dialog {
	height: 80%;
}

.modal-content {
	height: 50%;
	display: flex;
	flex-direction: column;
}

.modal-body {
	flex: 1;
	overflow-y: auto;
}
</style>