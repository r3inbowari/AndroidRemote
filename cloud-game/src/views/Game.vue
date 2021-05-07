<template>
  <div id="gameCp">
    <section id="body">
      <section class="section-title" :style="titleBackground">
        <div class="overlay">
          <div class="content">
            <ol class="breadcrumb">
              <li><a href="/" id="nav_home">主界面</a></li>
              <li><a href="/games/648">魔灵召唤：失落的世纪</a></li>
            </ol>
            <h2 class="title">魔灵召唤：失落的世纪</h2>
          </div>
        </div>
      </section>

      <div id="content">
        <section class="game-detail">
          <div class="content">
            <div class="box">
              <div class="video">
                <iframe
                  src="https://player.bilibili.com/player.html?cid=99269782&aid=56824180&page=1&as_wide=1&high_quality=1&danmaku=1&t=1"
                  allowfullscreen="allowfullscreen"
                  width="100%"
                  height="500"
                  scrolling="no"
                  frameborder="0"
                  sandbox="allow-top-navigation allow-same-origin allow-forms allow-scripts"
                ></iframe>
              </div>
              <div class="profile">
                <div
                  class="game_icon"
                  style="
                    background-image: url(https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1612349547427.jpg);
                  "
                ></div>
                <div class="game">
                  <div class="inner">
                    <h3 class="title">魔灵召唤：失落的世纪</h3>
                    <p class="genre">Strategy</p>
                  </div>
                </div>
                <ul class="introduce_txt">
                  <li><span>精彩的巅峰对决！</span></li>
                </ul>
                <ul class="component">
                  <li class="item">
                    <a
                      href="https://www.facebook.com/LostCenturia"
                      target="_blank"
                      ><span><i class="icon_community"></i>社区</span></a
                    >
                  </li>
                  <li class="item">
                    <button type="button">
                      <i class="icon_store"></i>开始游戏
                    </button>
                    <ul class="menu">
                      <li>
                        <a
                          href="http://play.withhive.com/r?c=11964"
                          target="_blank"
                        >
                          <i class="icon_appstore"></i>App Store
                        </a>
                      </li>
                      <li>
                        <a
                          href="http://play.withhive.com/r?c=11942"
                          target="_blank"
                        >
                          <i class="icon_googleplay"></i>Google Play Store
                        </a>
                      </li>
                    </ul>
                  </li>
                  <li class="item">
                    <button type="button">
                      <i class="icon_share"></i>分享
                    </button>
                    <ul class="menu">
                      <li class="btn_fb">
                        <a href="#" target="_blank"
                          ><i class="icon_fb"></i>脸书</a
                        >
                      </li>
                      <li class="btn_link">
                        <a href="javascript:void(0);"
                          ><i class="icon_copy"></i>复制链接</a
                        >
                      </li>
                    </ul>
                  </li>
                </ul>
              </div>
            </div>
            <!-- 说明 -->
            <div class="desc" id="description">
              <h4>游戏说明</h4>
              <div class="txt" style="height: auto">
                <div class="wrap">
                  <p>◈ [魔灵召唤：失落的世纪]精彩的巅峰对决！</p>
                  <p><br /></p>
                  <p>◇ 官方网页： https://summonerswar.com/LostCenturia</p>
                  <p>◇ 官方论坛： https://cafe.naver.com/lostcenturia</p>
                  <p>◇ 官方脸书：https://www.facebook.com/LostCenturia</p>
                  <p>
                    ◇
                    官方kakao频道：http://pf.kakao.com/_zxffvK&ZeroWidthSpace;<br />
                  </p>
                </div>
              </div>
              <button type="button" class="view_more" style="display: none">
                查看更多<i class="icon"></i>
              </button>
            </div>
            <div class="screenshot">
              <DetailSlider></DetailSlider>
            </div>
            <div class="desc" id="features">
              <h4>特征</h4>
              <div class="txt on">
                <div class="wrap">
                  <p>深受全球1亿召唤师喜爱的‘魔灵召唤’的另一个故事</p>
                  <p>
                    - 前往至[魔灵召唤：失落的世纪]，查看天空之役之前的故事吧！
                  </p>
                  <p><br /></p>
                  <p>■ 与全球召唤师展开实时对战吧！</p>
                  <p>- 华丽的动作与实时对战，令人窒息的PvP对决！</p>
                  <p>- 挑战成为全球第1名，展开世界级的对战吧。</p>
                  <p>- 战略战斗与动作，挑战最高等级的竞技场吧。</p>
                  <p><br /></p>
                  <p>■ 实时策略动作手游，以不可预测的反击技能压倒对手！</p>
                  <p>- 通过多种PvP与PvE模式，体验精彩的动作战略战斗！</p>
                  <p>- 通过‘反击技能’扭转局势，获得胜利吧！</p>
                  <p>- 使用失落的世纪独有的召唤师咒语技能吧。</p>
                  <p><br /></p>
                  <p>■ 创建联盟与好友构建交流！</p>
                  <p>- 向联盟成员及好友申请对战。</p>
                  <p>- 与联盟成员分享魔灵与咒语。</p>
                  <p><br /></p>
                  <p>■ 召唤个性突出的魔灵！</p>
                  <p>- 以失落的世纪风格，重新诠释天空之城的魔灵们。&nbsp;</p>
                  <p>- 用自己的策略，组合符文及技能石。</p>
                  <p>- 体验最强战略及战斗动作手游！&ZeroWidthSpace;<br /></p>
                </div>
              </div>
              <button type="button" class="view_more on">
                查看更多<i class="icon"></i>
              </button>
            </div>
          </div>
        </section>
      </div>

      <section class="other-games"></section>
    </section>
    <footer class="footer">
      <HomeFooter></HomeFooter>
    </footer>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  onMounted,
  getCurrentInstance,
  ref,
  reactive,
  watch,
  toRefs,
  onBeforeRouteUpdate,
} from 'vue'

import { useRoute } from 'vue-router'

import HomeFooter from '../components/Footer.vue'

import DetailSlider from '../components/DetailSlider.vue'

export default defineComponent({
  name: 'GameDetail',
  components: { HomeFooter, DetailSlider },
  props: {},
  setup() {
    const route = useRoute()

    let titleBackground = ref('')

    function changeBackground(r, url) {
      r.value = 'background-image: url("' + url + '")'
    }

    let aynx = {
      title: '内核活动2',
      alias: 'kernel event',
      cover:
        'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20210204/0aa3bbb693b4d1b9cb6fc61ce9d9ac8e_1200x490.jpg',
      icon:
        'https://image-glb.qpyou.cn/hubweb/gmnotice/appcenter/1612412918869.jpg',
    }

    onMounted(() => {
      changeBackground(
        titleBackground,
        'https://image-glb.qpyou.cn/hubweb/hive_img/web/banner/20210204/0aa3bbb693b4d1b9cb6fc61ce9d9ac8e_1200x490.jpg'
      )
    })
    return {
      titleBackground,
    }
  },
})
</script>

<style>
body {
  line-height: 1;
  background: #f7f8fa;
}

/* 屏蔽 */
#gameCp .footer .content {
  padding: 50px 20px;
}
</style>

<style scoped>
.section-title {
  padding: 0;
  background-color: rgba(10, 10, 18, 0.85);
  background-repeat: no-repeat;
  background-position: 50% 50%;
  background-size: cover;
  position: relative;
  margin: 0 auto;
}

* {
  margin: 0;
  padding: 0;
  vertical-align: baseline;
  border: 0;
  font-family: 'Noto Sans', sans-serif;
}

.section-title .overlay {
  padding: 44px 0;
  text-align: center;
  background-image: url(../assets/overlay_title_section.png);
  background-repeat: repeat-x;
  background-size: 1px 180px;
}

.content {
  width: 1240px;
  position: relative;
  margin: 0 auto;
  padding: 0 20px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

.game-detail {
  background-color: rgb(247, 248, 250);
  padding: 110px 0 60px;
}

.section-title .breadcrumb {
  color: #7e8592;
  font-family: 'Noto Sans SC', sans-serif;
}
.section-title .title {
  margin-top: 14px;
  font-size: 32px;
  line-height: 38px;
  color: #fff;
  font-family: 'Noto Sans SC', sans-serif;
}

ol,
ul {
  list-style: none;
}

ol {
  display: block;
  list-style-type: decimal;
  margin-block-start: 1em;
  margin-block-end: 1em;
  margin-inline-start: 0px;
  margin-inline-end: 0px;
  padding-inline-start: 40px;
}

a {
  color: inherit;
  text-decoration: none;
  outline: 0;
}

.section-title .breadcrumb li {
  display: inline;
  font-size: 14px;
  line-height: 20px;
}

.section-title .breadcrumb li + li:before {
  padding: 0 2px;
  color: #7e8592;
  content: '>';
  font-family: 'Noto Sans SC', sans-serif;
}

/* 游戏简介 */
.game-detail .profile {
  width: 1158px;
  height: 429px;
  margin-left: 42px;
  padding: 30px 20px 100px 753px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  background-color: #fff;
  border-radius: 20px;
}

.game-detail .profile .game_icon {
  width: 130px;
  height: 130px;
  float: left;
  top: 20px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  background-size: 100% 100%;
  border-radius: 30px;
}

.game-detail .profile .game {
  width: 215px;
  height: 130px;
  display: table;
  float: left;
  margin-bottom: 20px;
  padding: 0 20px;
}

.game-detail .profile .game .inner {
  display: table-cell;
  vertical-align: middle;
}

.section_game_detail .profile .introduce_txt {
  height: 290px;
}

.game-detail .profile .introduce_txt li:first-child {
  margin-top: 0;
}

.game-detail .profile .component {
  width: 385px;
  height: 72px;
  display: table;
  position: relative;
  left: -2px;
  margin: 0 0 0 0;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
  border-radius: 12px;
}

.game-detail .profile .component .item:first-child {
  border-left: 2px solid #e7e9ef;
  border-top-left-radius: 12px;
  border-bottom-left-radius: 12px;
}

.game-detail .profile .component .item {
  width: 33.33%;
  display: table-cell;
  position: relative;
  overflow: hidden;
  box-sizing: border-box;
  border-top: 2px solid #e7e9ef;
  border-bottom: 2px solid #e7e9ef;
}

.game-detail .profile .component .item:last-child {
  border-right: 2px solid #e7e9ef;
  border-top-right-radius: 12px;
  border-bottom-right-radius: 12px;
}

.game-detail .profile .component .icon_store {
  background-position: -459px -148px;
}

.game-detail .profile .component [class^='icon_'] {
  width: 20px;
  height: 20px;
  display: block;
  position: relative;
  left: 50%;
  margin: 0 0 4px -9px;
  vertical-align: middle;
}

.game-detail .profile .component .item .menu {
  width: 218px;
  position: absolute;
  z-index: -1;
  top: 74px;
  left: 50%;
  margin-left: -109px;
  padding: 12px 0;
  color: #fff;
  background-color: #000;
  border-radius: 12px;
}

.game-detail .profile .component .item .menu:before {
  width: 16px;
  height: 10px;
  position: absolute;
  top: -10px;
  left: 50%;
  margin-left: -8px;
  content: '';
  background-position: -184px -175px;
}
.game-detail .profile .component:after {
  display: table;
  clear: both;
  content: '';
}

.icon,
.arrow,
.section_notice_detail .notice_article .share button,
.section_event_list .list .event.fixed:after,
.search_box .btn_input_delete,
.section_customer_service .menu [class^='icon'],
.pop_movie .wrap .btn_close:after,
.section_notice_list .board .board_list li.fixed a:before,
.game-detail .profile .component [class^='icon_'],
.game-detail .profile .component .item .menu:before,
.game-detail .profile .component .item .menu:after,
.game-detail .profile .component .item .menu:before,
.pop_link_copy .wrap .btn_close:after,
.pop_link_copy .icon_copy,
.pop_qrcode .wrap .btn_close:after,
.toast_alert .txt:before,
.toast_alert .btn_close .icon,
.pop_link_copy .txt:before,
.pop_link_copy .btn_close .icon,
.section_electric_ja .list li:before,
.game-detail .video.external .btn_external {
  background-image: url(../assets/icon_2x.png);
  background-repeat: no-repeat;
  background-size: 512px 512px;
}

.game-detail .profile .component .icon_community {
  background-position: -439px -148px;
}

.game-detail .profile .component .item:after {
  width: 1px;
  height: 44px;
  position: absolute;
  top: 50%;
  right: 0;
  margin-top: -22px;
  content: '';
  background-color: #e7e9ef;
}

.game-detail .profile .component .item > a > span {
  display: inline-block;
  vertical-align: middle;
}

.game-detail .profile .component .icon_share {
  background-position: -478px -148px;
}

.game-detail .profile .component .item > a,
.game-detail .profile .component button {
  width: 100%;
  height: 100%;
  display: block;
  position: absolute;
  font-size: 14px;
  font-weight: bold;
  color: #7e8592;
  text-align: center;
}

.game-detail .profile .component .item > a:after {
  width: 0;
  height: 100%;
  display: inline-block;
  vertical-align: middle;
  content: '';
}

button {
  background-color: white;
}

/* 字体 */
.game-detail .profile .introduce_txt {
  height: 290px;
}

.game-detail .profile .introduce_txt li {
  width: 100%;
  margin-top: 10px;
  font-size: 16px;
  font-weight: bold;
  line-height: 24px;
  color: #7e8692;
  font-family: 'Noto Sans SC', sans-serif;
}

.game-detail .profile .game .title {
  margin-bottom: 10px;
  font-size: 24px;
  font-weight: bold;
  line-height: 30px;
  color: #000f45;
}

.game-detail .profile .game .genre {
  font-size: 14px;
  line-height: 20px;
  color: #1277ff;
}

.game-detail .video {
  width: 763px;
  height: 429px;
  position: absolute;
  overflow: hidden;
  top: -40px;
  left: 20px;
  background-color: #000;
  background-size: 100% 100%;
  border-radius: 8px;
  box-shadow: 0 13px 20px -6px #aab3c5;
}

.game-detail .desc {
  width: 100%;
  position: relative;
  padding: 60px 0 0 42px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

/* 说明 */
.game-detail .desc {
  width: 100%;
  position: relative;
  padding: 60px 0 0 42px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

.game-detail .desc .txt {
  width: 100%;
  height: 132px;
  overflow: hidden;
  margin-bottom: 40px;
  font-size: 15px;
  line-height: 22px;
  color: #7e8592;
}

.wrap {
  min-width: 1240px;
  overflow: hidden;
  background-color: #f7f8fa;
}

.game-detail .desc .view_more {
  font-size: 14px;
  font-weight: bold;
  line-height: 20px;
  color: #1277ff;
}

.game-detail .desc .view_more .icon {
  width: 11px;
  height: 6px;
  display: inline-block;
  position: relative;
  margin: -2px 0 0 10px;
  vertical-align: middle;
  background-position: -433px -86px;
}

.game-detail .desc h4 {
  padding-bottom: 20px;
  font-size: 22px;
  font-weight: bold;
  line-height: 28px;
  color: #000f45;
}

/* 描述 */
.game-detail .screenshot + .desc {
  padding: 0 0 60px 42px;
}

.game-detail .desc {
  width: 100%;
  position: relative;
  padding: 60px 0 0 42px;
  -webkit-box-sizing: border-box;
  box-sizing: border-box;
}

.game-detail .desc .txt.on {
  height: auto;
}

.game-detail .desc .view_more.on {
  display: none;
}

.footer {
  background-color: #0a0a12;
}

.game-detail .screenshot {
  position: relative;
  padding: 60px 0 30px;
}
</style>
