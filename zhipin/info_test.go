package zhipin

import "testing"

func TestParsePosition(t *testing.T) {
	html := `<div id="main">
        <div class="job-banner">
            <div class="inner home-inner">
                <div class="job-primary detail-box">
                    <div class="pos-bread">
                        <a href="/" ka="job-breadcrumb-top1">首页</a>
                        <i class="icon-arrow-right"></i>
                        <a href="/c101010100/" ka="job-breadcrumb-top2">北京招聘</a>·<a href="/p100116/" ka="job-breadcrumb-top21">Golang招聘</a>
                        <i class="icon-arrow-right"></i>
                        <a href="/c101010100-p100116/" ka="job-breadcrumb-top3">北京Golang招聘</a>
                        <i class="icon-arrow-right"></i>
                        <a href="/job_detail/156cdef2df73d1391Xxy2ti0GFc~.html" ka="job-breadcrumb-top4">小米Golang开发工程师招聘</a>
                    </div>
                    <div class="info-primary">
                        <div class="job-status">招聘中</div>
                        <div class="name">
                            <h1>Golang开发工程师</h1>
                            <span class="salary">
                                    18k-35k
                            </span>
                        </div>
                        <p>北京<em class="dolt"></em>3-5年<em class="dolt"></em>大专</p>
                            <div class="tag-container">
                                <div class="tag-more" style="opacity: 1; display: block; left: 457px;">
                                    <span class="link-more">...</span>
                                    <div class="tag-all job-tags">
                                        <span>五险一金</span><span>补充医疗保险</span><span>定期体检</span><span>年终奖</span><span>股票期权</span><span>带薪年假</span><span>员工旅游</span><span>免费班车</span><span>餐补</span><span>交通补助</span><span>节日福利</span><span>住房补贴</span>
                                    </div>
                                </div>
                                <div class="job-tags">
                                    <span>五险一金</span><span>补充医疗保险</span><span>定期体检</span><span>年终奖</span><span>股票期权</span><span>带薪年假</span><span>员工旅游</span><span>免费班车</span><span>餐补</span><span>交通补助</span><span>节日福利</span><span>住房补贴</span>
                                </div>
                            </div>
                    </div>
                    <div class="job-op">
                        <div class="btn-container">
                            <!-- 未登录 -->
                                                 <a class="btn btn-startchat" ka="go_greet_tosign_28815985" href="javascript:;" redirect-url="/geek/new/index/chat?id=d8d31f652810f7121nZ72t6-GFA~" target="_blank" data-url="/gchat/addRelation.json?jobId=156cdef2df73d1391Xxy2ti0GFc~&amp;lid=1r39vKT8qKM.search">立即沟通</a>

                        </div>
                        <div class="op-container">
                            <!-- 在线简历 -->
                                <a href="/geek/attresume/parser.html" class="fl icon-resume" ka="job_online_resume_0"><i class="icon"></i>填写在线简历</a>
                            <!-- 附件简历 -->
                                <a href="/geek/attresume/parser.html" class="fr icon-upload" ka="job_attach_resume_0"><i class="icon"></i>上传附件简历</a>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="smallbanner" style="display: none;">
            <div class="inner home-inner">
                <div class="detail-op">
                    <div class="btns">
                                             <a class="btn btn-startchat" ka="go_greet_tosign_28815985" href="javascript:;" redirect-url="/geek/new/index/chat?id=d8d31f652810f7121nZ72t6-GFA~" target="_blank" data-url="/gchat/addRelation.json?jobId=156cdef2df73d1391Xxy2ti0GFc~&amp;lid=1r39vKT8qKM.search">立即沟通</a>

                    </div>
                </div>
                <div class="company-info">
                    <div class="name"><h1>Golang开发工程师</h1> <span class="badge">18k-35k</span></div>
                        <div class="tag-container">
                            <div class="tag-more" style="opacity: 1; display: block; left: 457px;">
                                <span class="link-more">...</span>
                                <div class="tag-all job-tags">
                                    <span>五险一金</span><span>补充医疗保险</span><span>定期体检</span><span>年终奖</span><span>股票期权</span><span>带薪年假</span><span>员工旅游</span><span>免费班车</span><span>餐补</span><span>交通补助</span><span>节日福利</span><span>住房补贴</span>
                                </div>
                            </div>
                            <div class="job-tags">
                                <span>五险一金</span><span>补充医疗保险</span><span>定期体检</span><span>年终奖</span><span>股票期权</span><span>带薪年假</span><span>员工旅游</span><span>免费班车</span><span>餐补</span><span>交通补助</span><span>节日福利</span><span>住房补贴</span>
                            </div>
                        </div>
                    <div class="info">
                                小米
                        <a href="/gongsir/6f1aa1d6b1d033ad33B43N0~.html" target="_blank" class="link-more" ka="look-all-position">查看所有职位</a>
                    </div>
                </div>
            </div>
        </div>
        <div class="job-box">
            <div class="inner home-inner">
                <div class="job-sider">

    <div class="sign-wrap sign-wrap-v2">
        <div class="sign-form sign-quick" style="display: block;">
            <canvas id="sign-side" width="298" height="50"></canvas>
            <h3 class="title">各大行业职位任你选</h3>
            <form action="/wapi/zppassport/login/phone" method="post" ka="sup_g_right_done"><input type="hidden" name="pk" value="cpc_job_detail">
                <div class="tip-error tip-error-form"></div>
                <div class="form-row row-select">
                    <span class="dropdown-select"><i class="icon-select-arrow"></i><em class="text-select">+86</em><input type="hidden" name="regionCode" value="+86"></span>
                    <span class="ipt-wrap"><i class="icon-sign-phone"></i><input type="tel" class="ipt ipt-phone required" placeholder="手机号" name="phone"></span>
                    <div class="dropdown-menu">
<ul>
    <li data-val="+86"><span class="num">+86</span>中国大陆</li>
    <li data-val="+1"><span class="num">+1</span>美国</li>
    <li data-val="+852"><span class="num">+852</span>香港</li>
    <li data-val="+81"><span class="num">+81</span>日本</li>
    <li data-val="+886"><span class="num">+886</span>台湾</li>
    <li data-val="+44"><span class="num">+44</span>英国</li>
    <li data-val="+82"><span class="num">+82</span>韩国</li>
    <li data-val="+33"><span class="num">+33</span>法国</li>
    <li data-val="+7"><span class="num">+7</span>俄罗斯</li>
    <li data-val="+39"><span class="num">+39</span>意大利</li>
    <li data-val="+65"><span class="num">+65</span>新加坡</li>
    <li data-val="+63"><span class="num">+63</span>菲律宾</li>
    <li data-val="+60"><span class="num">+60</span>马来西亚</li>
    <li data-val="+64"><span class="num">+64</span>新西兰</li>
    <li data-val="+34"><span class="num">+34</span>西班牙</li>
    <li data-val="+95"><span class="num">+95</span>缅甸</li>
    <li data-val="+230"><span class="num">+230</span>毛里求斯</li>
    <li data-val="+263"><span class="num">+263</span>津巴布韦</li>
    <li data-val="+20"><span class="num">+20</span>埃及</li>
    <li data-val="+92"><span class="num">+92</span>巴基斯坦</li>
    <li data-val="+880"><span class="num">+880</span>孟加拉国</li>
</ul>                    </div>
                    <div class="tip-error"></div>
                </div>
                <div class="form-row">
                    <div class="row-code nc-container" id="verrifyn8joi8" data-nc-idx="3">
<div id="nc_3_wrapper" class="nc_wrapper">
<div id="nc_3_n1t" class="nc_scale">
<div id="nc_3__bg" class="nc_bg"></div>
<span id="nc_3_n1z" class="nc_iconfont btn_slide"></span>
<div id="nc_3__scale_text" class="scale_text slidetounlock"><span class="nc-lang-cnt" data-nc-lang="_startTEXT">请按住滑块，拖动到最右边</span></div>
<div id="nc_3_clickCaptcha" class="clickCaptcha">
<div class="clickCaptcha_text">
<b id="nc_3__captcha_text" class="nc_captch_text"></b>
<i id="nc_3__btn_2" class="nc_iconfont nc_btn_2 btn_refresh"></i>
</div>
<div class="clickCaptcha_img"></div>
<div class="clickCaptcha_btn"></div>
</div>
<div id="nc_3_imgCaptcha" class="imgCaptcha">
<div class="imgCaptcha_text"><input id="nc_3_captcha_input" maxlength="6" type="text" style="ime-mode:disabled"></div>
<div class="imgCaptcha_img" id="nc_3__imgCaptcha_img"></div>
<i id="nc_3__btn_1" class="nc_iconfont nc_btn_1 btn_refresh" onclick="document.getElementById('nc_3__imgCaptcha_img').children[0].click()"></i>
<div class="imgCaptcha_btn">
<div id="nc_3__captcha_img_text" class="nc_captcha_img_text"></div>
<div id="nc_3_scale_submit" class="nc_scale_submit"></div>
</div>
</div>
<div id="nc_3_cc" class="nc-cc"></div>
<i id="nc_3__voicebtn" tabindex="0" role="button" class="nc_voicebtn nc_iconfont" style="display:none"></i>
<b id="nc_3__helpbtn" class="nc_helpbtn"><span class="nc-lang-cnt" data-nc-lang="_learning">了解新功能</span></b>
</div>
<div id="nc_3__voice" class="nc_voice"></div>
</div>
</div>
                    <div class="tip-error"></div>
                </div>
                <input type="hidden" name="csessionId" value="">
                <input type="hidden" name="csig" value="">
                <input type="hidden" name="ctoken" value="">
                <input type="hidden" name="cscene" value="nc_login">
                <input type="hidden" value="FFFF0N00000000006DC1" name="cappKey">
                <div class="form-row">
                    <span class="ipt-wrap">
                        <i class="icon-sign-sms"></i>
                        <input type="text" class="ipt ipt-sms required" ka="signup-sms" placeholder="短信验证码" name="phoneCode" maxlength="6">
                        <input type="hidden" name="smsType" value="7">
                        <button type="button" class="btn btn-sms" data-url="/wapi/zppassport/send/smsCode">发送验证码</button>
                    </span>
                    <div class="tip-error"></div>
                </div>
                <div class="form-btn">
                    <button type="submit" class="btn" ka="sup_g_right">登录/注册</button>
                </div>
                <div class="text-tip">
                    <input type="checkbox" class="agree-policy" name="policy">我已同意<a href="https://www.zhipin.com/register/protocol/introduce" target="_blank">用户协议及隐私政策</a>
                </div>
            </form>
        </div>
    </div>

                    <div class="sider-company">
                        <p class="title">公司基本信息</p>
                            <div class="company-info">
                                <a ka="job-detail-company" href="/gongsi/6f1aa1d6b1d033ad33B43N0~.html" title="                                    小米
" target="_blank">
                                        <img src="https://img.bosszhipin.com/beijin/mcs/bar/brand/84270.jpg?x-oss-process=image/resize,w_120,limit_0" alt="">
                                </a>
                                <a ka="job-detail-company" href="/gongsir/6f1aa1d6b1d033ad33B43N0~.html" title="                                    小米
招聘" target="_blank">
                                                                        小米

                                </a>
                            </div>
                                 <p><i class="icon-stage"></i>已上市</p>
                                <p><i class="icon-scale"></i>10000人以上</p>
                                <p><i class="icon-industry"></i><a ka="job-detail-brandindustry" href="/i100020/">互联网</a></p>
                                <p><i class="icon-net"></i>http://www.mi.com</p>
                    
                </div>

                    <div class="salary-miniapp">
                        <img src="https://img.bosszhipin.com/beijin/app/mobile/WX20181115-144232-4ed035aa4fa356bf0625f088b6f5b21e.png">
                    </div>
                <div class="five-graph-test promotion-img">
                    <h4>五维职业性格测评</h4>
                    <i></i>
                    <p>专业分析职场优势</p>
                    <a class="btn" href="/activity/personality/index.html" target="_blank" ka="job-detail-personality">立即测试</a>
                </div>
                    <div class="promotion-job">
                        <h3><a href="/c101010100-p100116/e_105-y_5_6_7/" ka="more-similar-jobs2" class="more" target="_blank">更多相似职位</a>相似职位</h3>
                        <ul>
                                <li>
                                    <div class="company-logo">
                                        <a ka="job_sug_brand1_custompage" href="/gongsi/2a56a20c9b7ff2251nJ83tm5.html" target="_blank" title="甘来信息科技">
                                                <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20170602/18803c15b3f8c22913474c31207467d2706d1c97828cd901df36ea269a340040.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                        </a>
                                    </div>
                                    <div class="info-primary">
                                        <div class="name"><a ka="job_sug_1" href="/job_detail/19ca60bf3f24e77a1nFz39q7GFY~.html" target="_blank">Go开发工程师 <span class="red">20k-40k</span></a></div>
                                        <p class="gray"><a ka="job_sug_brandjob1_custompage" href="/gongsir/2a56a20c9b7ff2251nJ83tm5.html" target="_blank" title="甘来信息科技招聘">甘来信息科技</a><em class="vdot">·</em>北京</p>
                                    </div>
                                </li>
                                <li>
                                    <div class="company-logo">
                                        <a ka="job_sug_brand2_custompage" href="/gongsi/941568145abcddfd3nJ43to~.html" target="_blank" title="淡蓝网Blued">
                                                <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20170609/2d55ebc7522a8dddd3a4609aabc070faf91b2c385ac1d6d895e0980a4016c3e3.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                        </a>
                                    </div>
                                    <div class="info-primary">
                                        <div class="name"><a ka="job_sug_2" href="/job_detail/8c800b0dc9c315471XJ73Nu6EVY~.html" target="_blank">golang开发工程师 <span class="red">20k-35k</span></a></div>
                                        <p class="gray"><a ka="job_sug_brandjob2_custompage" href="/gongsir/941568145abcddfd3nJ43to~.html" target="_blank" title="淡蓝网Blued招聘">淡蓝网Blued</a><em class="vdot">·</em>北京</p>
                                    </div>
                                </li>
                                <li>
                                    <div class="company-logo">
                                        <a ka="job_sug_brand3_custompage" href="/gongsi/725311d303ead3261nx40tW7FA~~.html" target="_blank" title="神话传媒合肥分公司">
                                                <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20180511/c09779fd753eec27dd612f4a1168e950cfcd208495d565ef66e7dff9f98764da.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                        </a>
                                    </div>
                                    <div class="info-primary">
                                        <div class="name"><a ka="job_sug_3" href="/job_detail/0b8410f9149bcbe01Xd_3t25EVE~.html" target="_blank">golang开发工程师 <span class="red">26k-52k</span></a></div>
                                        <p class="gray"><a ka="job_sug_brandjob3_custompage" href="/gongsir/725311d303ead3261nx40tW7FA~~.html" target="_blank" title="神话传媒合肥分公司招聘">神话传媒合肥分公司</a><em class="vdot">·</em>北京</p>
                                    </div>
                                </li>
                                <li>
                                    <div class="company-logo">
                                        <a ka="job_sug_brand4_custompage" href="/gongsi/546d2237323ed5bd1Xd92NW4.html" target="_blank" title="SmartX">
                                                <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20170105/6474e8e06e086e2abb97eae0cc812b8347f70f7a8b3e657bb655f8122951552e.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                        </a>
                                    </div>
                                    <div class="info-primary">
                                        <div class="name"><a ka="job_sug_4" href="/job_detail/11e018fe859162681XF62tu4FlY~.html" target="_blank">后台开发工程师 （Python、Golang） <span class="red">20k-40k</span></a></div>
                                        <p class="gray"><a ka="job_sug_brandjob4_custompage" href="/gongsir/546d2237323ed5bd1Xd92NW4.html" target="_blank" title="SmartX招聘">SmartX</a><em class="vdot">·</em>北京</p>
                                    </div>
                                </li>
                                <li>
                                    <div class="company-logo">
                                        <a ka="job_sug_brand5_custompage" href="/gongsi/980f48937a13792b1nd63d0~.html" target="_blank" title="滴滴出行">
                                                <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20190408/c23f08b24983fffa26a3a8ba19a463523cc05a6873981b80bf124ddd6c45f629_s.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                        </a>
                                    </div>
                                    <div class="info-primary">
                                        <div class="name"><a ka="job_sug_5" href="/job_detail/c954f89cb6da25830n1-2N2-Fw~~.html" target="_blank">Golang开发工程师 <span class="red">25k-50k</span></a></div>
                                        <p class="gray"><a ka="job_sug_brandjob5_custompage" href="/gongsir/980f48937a13792b1nd63d0~.html" target="_blank" title="滴滴出行招聘">滴滴出行</a><em class="vdot">·</em>北京</p>
                                    </div>
                                </li>
                        </ul>
                            <div class="view-more"><a href="/c101010100-p100116/e_105-y_5_6_7/" target="_blank" ka="more-similar-jobs3"><i class="more">更多相似职位</i></a></div>
                    </div>

                <div class="promotion-img">
                    <a href="/app.html" ka="job-detail-app" target="_blank"><img src="https://static.zhipin.com/zhipin/v147/web/geek/images/pro-1.png" alt=""></a>
                </div>
            </div>

            <div class="job-detail">
                <div class="detail-op">
                    <div class="btns">
                        <div class="op-links">
                            <a ka="job_detail_wechat_share" href="javascript:;" class="link-wechat-share" data-url="/boss/share/job.json?jobId=156cdef2df73d1391Xxy2ti0GFc~&amp;source=2">微信扫码分享</a>
                            <a href="javascript:;" ka="job_detail_like" class="link-like " data-url="/geek/tag/jobtagupdate.json?jobId=156cdef2df73d1391Xxy2ti0GFc~&amp;expectId=&amp;tag=4&amp;lid=1r39vKT8qKM.search">感兴趣</a>
                            <a ka="job_detail_report" href="javascript:;" class="link-report">举报</a>
                        </div>
                    </div>

                    <div class="detail-figure">
                            <img src="https://img2.bosszhipin.com/boss/avatar/avatar_11.png?x-oss-process=image/resize,w_100,limit_0" alt="">
                    </div>
                    <h2 class="name">刘女士<i class="icon-vip"></i></h2>
                    <p class="gray">HR<em class="vdot">·</em>刚刚在线</p>
                </div>
                <div class="detail-content">
                        <div class="job-sec">
                            <h3>职位描述</h3>
                            <div class="text">
                                工作职责：<br>- 负责小米商城后端研发工作（包括PC，移动端，小程序等），为小米用户提供高效可靠服务<br>- 与商城产品，市场活动策划进行密切沟通，通过技术手段，支持小米各项促销大活动<br>- 与团队一起研究解决电商中高并发，高可用系统，微服务架构，多机房，高可用缓存系统等技术难题<br>- 负责新技术研究，通过大数据和AI技术，帮助业务提升业绩，达成目标<br>工作要求：<br>- 计算机或相关专业本科以上学历，硕士学历加分<br>- 熟练掌握Go语言或PHP/JAVA 愿意转go语言，编程技术扎实，有相关Go项目经验优先<br>- 熟悉业务开发相关技术：常见数据库和协议，架构，存储，缓存，消息队列，API设计理念等<br>- 有相关互联网公司工作经验，有大流量项目经验优先<br>- 有良好的团队合作精神，能与产品、策划，测试等其他同事协作解决问题<br>- 对技术有热情，愿意挑战技术难题，业务难题<br>- 扎实的编程能力, 熟悉算法和数据结构, 熟悉计算机基础理论, 熟悉分布式系统原理
                            </div>
                        </div>
                            <div class="job-sec">
                                <h3>团队介绍</h3>
                                    <div class="text">可年后入职。</div>
                                    <div class="job-tags">
                                            <span>领导nice</span>
                                            <span>不打卡</span>
                                            <span>扁平管理</span>
                                            <span>公司氛围好</span>
                                            <span>年终分红</span>
                                            <span>免费零食</span>
                                            <span>股票期权</span>
                                            <span>带薪年假</span>
                                    </div>
                            </div>
                            <div class="job-sec company-info">
                                <h3>公司介绍</h3>
                                <div class="text">
                                    小米公司成立于2010年4月，是一家有实体经济的互联网公司。小米公司是专注于智能手机、智能家居、互联网电视的创新型科技企业。同时，在互联网金融（支付、信贷、保险、理财等）、互娱和影业等领域也积极布局，并初具规模。小米用互联网开发模式、极客精神研发产品，致力于让每个人都能享受科技的乐趣。<br>2011年
                                </div>
                                <a ka="job-comintroduce" href="/gongsi/6f1aa1d6b1d033ad33B43N0~.html" target="_blank" class="look-all">查看全部</a>
                            </div>
                        <div class="job-sec prop-item">
                            <h3>竞争力分析</h3>
                            <div class="">
                                <div class="title">综合竞争力评估</div>
                                <div class="prop-container"><p class="text-position">你在？位置</p>
                                    <span class="level-1"></span>
                                    <span class="level-2"><em>一般</em></span>
                                    <span class="level-3"><em>良好</em></span>
                                    <span class="level-4"><em>优秀</em></span>
                                    <span class="level-5"><em>极好</em></span>
                                </div>
                                <div class="prop-detail">
                                    <div class="pull-right"><a href="javascript:;" class="link-detail" ka="check_personal_competitive_detail">查看完整个人竞争力</a></div>
                                    <div>个人综合排名：<span>在<img src="https://static.zhipin.com/zhipin/v147/web/geek/images/pic-blur.png">人中排名第<img src="https://static.zhipin.com/zhipin/v147/web/geek/images/pic-blur.png"></span></div>
                                </div>
                            </div>
                        </div>
                            <div class="job-sec">
                                <h3>工商信息</h3>
                                <div class="name">小米科技有限责任公司</div>
                                <div class="level-list">
                                    <li><span>法人代表：</span>雷军</li>
                                    <li><span>注册资金：</span>185000万元人民币</li>
                                    <li class="res-time"><span>成立时间：</span>2010-03-03</li>
                                    <li class="company-type"><span>企业类型：</span>有限责任公司（自然人投资或控股）</li>
                                    <li class="manage-state"><span>经营状态：</span>开业</li>
                                </div>
                                <a ka="job-cominfo" href="/gongsi/6f1aa1d6b1d033ad33B43N0~.html" target="_blank" class="look-all">查看全部</a>
                            </div>
                        <div class="job-sec">
                            <h3>工作地址</h3>
                            <div class="job-location">
                                <div class="location-address">北京市 海淀区 小米总参</div>
                                <div class="job-location-map js-open-map" data-content="北京市 海淀区 小米总参">
                                    <img src="https://restapi.amap.com/v3/staticmap?size=652*174&amp;markers=mid,0xFF0000,A:116.337324,40.026647&amp;key=21b56a6cc83fad7668dbb0e9564759a7" alt="公司地址">
                                    <p>点击查看地图</p>
                                </div>
                            </div>
                        </div>
                        <div class="job-recomend">
                            <h3><a href="/c101010100-p100116/e_105-y_5_6_7/" ka="more-similar-jobs4" class="more" target="_blank">更多职位</a>看过该职位的人还看了</h3>
                            <ul>
                                    <li>
                                        <a ka="job_sug_6" href="/job_detail/1de28e3a99e011561HR_0tu1GFQ~.html" target="_blank">
                                            <div class="company-logo">
                                                    <img src="https://img.bosszhipin.com/beijin/mcs/chatphoto/20180507/56c4809c1a3637cb8ff850feb13aa82ecfcd208495d565ef66e7dff9f98764da.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                            </div>
                                            <div class="info-primary">
                                                <div class="name"><b>go开发工程师</b></div>
                                                <p class="red">20k-30k</p>
                                                <p class="gray">
                                                        波场<em class="vdot">·</em>
                                                    北京
                                                </p>
                                            </div>
                                        </a>
                                    </li>
                                    <li>
                                        <a ka="job_sug_7" href="/job_detail/dfcc14d32dd65bb71Hd72dy6FFI~.html" target="_blank">
                                            <div class="company-logo">
                                                    <img src="https://img.bosszhipin.com/beijin/mcs/bar/20180419/127a067dea373539c50d2875bce0ab95be1bd4a3bd2a63f070bdbdada9aad826.png?x-oss-process=image/resize,w_100,limit_0" alt="">
                                            </div>
                                            <div class="info-primary">
                                                <div class="name"><b>高级GO开发工程师</b></div>
                                                <p class="red">25k-50k</p>
                                                <p class="gray">
                                                        小年糕<em class="vdot">·</em>
                                                    北京
                                                </p>
                                            </div>
                                        </a>
                                    </li>
                                    <li>
                                        <a ka="job_sug_8" href="/job_detail/09da888788e0ed901XB909i0E1c~.html" target="_blank">
                                            <div class="company-logo">
                                                    <img src="https://img.bosszhipin.com/beijin/mcs/bar/brand/84270.jpg?x-oss-process=image/resize,w_100,limit_0" alt="">
                                            </div>
                                            <div class="info-primary">
                                                <div class="name"><b>go开发工程师</b></div>
                                                <p class="red">25k-35k</p>
                                                <p class="gray">
                                                        小米<em class="vdot">·</em>
                                                    北京
                                                </p>
                                            </div>
                                        </a>
                                    </li>
                            </ul>
                        </div>
                    <div class="search-box detail-search">
                        <div class="search-form">
                            <form action="/job_detail/" method="get" target="_blank">
                                <div class="search-form-con">
                                    <p class="ipt-wrap"><input ka="search-job-query" name="query" type="text" class="ipt-search" autocomplete="off" placeholder="搜索职位、公司" value="Golang"></p>
                                    <div class="city-sel" ka="search-select-city">
                                        <i class="line"></i>
                                        <span class="label-text cur"><b data-val="101010100">北京</b><i class="icon-arrow-down"></i></span>
                                    </div>
                                </div>
                                <input type="hidden" name="city" class="city-code" value="101010100">
                                <input type="hidden" name="source" value="9">
                                <button type="submit" ka="search-job" class="btn btn-search">搜索</button>
                                <div class="suggest-result">
                                    <ul><li>go</li><li>web</li><li>前端</li></ul>
                                </div>
                                <div class="city-box">
                                    <ul class="dorpdown-province"><li class="">热门</li><li ka="sel-province-1" class="">北京</li><li ka="sel-province-2" class="">上海</li><li ka="sel-province-3" class="">天津</li><li ka="sel-province-4" class="">重庆</li><li ka="sel-province-5" class="">黑龙江</li><li ka="sel-province-6" class="">吉林</li><li ka="sel-province-7" class="">辽宁</li><li ka="sel-province-8" class="">内蒙古</li><li ka="sel-province-9" class="">河北</li><li ka="sel-province-10" class="">山西</li><li ka="sel-province-11" class="">陕西</li><li ka="sel-province-12" class="">山东</li><li ka="sel-province-13" class="">新疆</li><li ka="sel-province-14" class="">青海</li><li ka="sel-province-15" class="">甘肃</li><li ka="sel-province-16" class="">宁夏</li><li ka="sel-province-17" class="">河南</li><li ka="sel-province-18" class="">江苏</li><li ka="sel-province-19" class="">湖北</li><li ka="sel-province-20" class="">浙江</li><li ka="sel-province-21" class="">安徽</li><li ka="sel-province-22" class="">福建</li><li ka="sel-province-23" class="">江西</li><li ka="sel-province-24" class="">湖南</li><li ka="sel-province-25" class="">贵州</li><li ka="sel-province-26" class="">四川</li><li ka="sel-province-27" class="">广东</li><li ka="sel-province-28" class="">云南</li><li ka="sel-province-29" class="">广西</li><li ka="sel-province-30" class="">海南</li><li ka="sel-province-31" class="">台湾</li><li ka="sel-province-32" class="">西藏</li><li ka="sel-province-33" class="">香港</li><li ka="sel-province-34" class="">澳门</li></ul>
                                    <div class="dorpdown-city"><ul><li ka="hot-city-100010000" data-val="100010000" class="cur">全国</li><li ka="hot-city-101010100" data-val="101010100" class="cur">北京</li><li ka="hot-city-101020100" data-val="101020100" class="cur">上海</li><li ka="hot-city-101280100" data-val="101280100" class="cur">广州</li><li ka="hot-city-101280600" data-val="101280600" class="cur">深圳</li><li ka="hot-city-101210100" data-val="101210100" class="cur">杭州</li><li ka="hot-city-101030100" data-val="101030100" class="cur">天津</li><li ka="hot-city-101110100" data-val="101110100" class="cur">西安</li><li ka="hot-city-101190400" data-val="101190400" class="cur">苏州</li><li ka="hot-city-101200100" data-val="101200100" class="cur">武汉</li><li ka="hot-city-101230200" data-val="101230200" class="cur">厦门</li><li ka="hot-city-101250100" data-val="101250100" class="cur">长沙</li><li ka="hot-city-101270100" data-val="101270100" class="cur">成都</li><li ka="hot-city-101180100" data-val="101180100" class="cur">郑州</li><li ka="hot-city-101040100" data-val="101040100" class="cur">重庆</li></ul><ul><li ka="hot-city-101010100" data-val="101010100" class="cur">北京</li></ul><ul><li ka="hot-city-101020100" data-val="101020100" class="cur">上海</li></ul><ul><li ka="hot-city-101030100" data-val="101030100" class="cur">天津</li></ul><ul><li ka="hot-city-101040100" data-val="101040100" class="cur">重庆</li></ul><ul><li ka="hot-city-101050100" data-val="101050100" class="cur">哈尔滨</li><li ka="hot-city-101050200" data-val="101050200" class="cur">齐齐哈尔</li><li ka="hot-city-101050300" data-val="101050300" class="cur">牡丹江</li><li ka="hot-city-101050400" data-val="101050400" class="cur">佳木斯</li><li ka="hot-city-101050500" data-val="101050500" class="cur">绥化</li><li ka="hot-city-101050600" data-val="101050600" class="cur">黑河</li><li ka="hot-city-101050700" data-val="101050700" class="cur">伊春</li><li ka="hot-city-101050800" data-val="101050800" class="cur">大庆</li><li ka="hot-city-101050900" data-val="101050900" class="cur">七台河</li><li ka="hot-city-101051000" data-val="101051000" class="cur">鸡西</li><li ka="hot-city-101051100" data-val="101051100" class="cur">鹤岗</li><li ka="hot-city-101051200" data-val="101051200" class="cur">双鸭山</li><li ka="hot-city-101051300" data-val="101051300" class="cur">大兴安岭</li></ul><ul><li ka="hot-city-101060100" data-val="101060100" class="cur">长春</li><li ka="hot-city-101060200" data-val="101060200" class="cur">吉林</li><li ka="hot-city-101060300" data-val="101060300" class="cur">四平</li><li ka="hot-city-101060400" data-val="101060400" class="cur">通化</li><li ka="hot-city-101060500" data-val="101060500" class="cur">白城</li><li ka="hot-city-101060600" data-val="101060600" class="cur">辽源</li><li ka="hot-city-101060700" data-val="101060700" class="cur">松原</li><li ka="hot-city-101060800" data-val="101060800" class="cur">白山</li><li ka="hot-city-101060900" data-val="101060900" class="cur">延边</li></ul><ul><li ka="hot-city-101070100" data-val="101070100" class="cur">沈阳</li><li ka="hot-city-101070200" data-val="101070200" class="cur">大连</li><li ka="hot-city-101070300" data-val="101070300" class="cur">鞍山</li><li ka="hot-city-101070400" data-val="101070400" class="cur">抚顺</li><li ka="hot-city-101070500" data-val="101070500" class="cur">本溪</li><li ka="hot-city-101070600" data-val="101070600" class="cur">丹东</li><li ka="hot-city-101070700" data-val="101070700" class="cur">锦州</li><li ka="hot-city-101070800" data-val="101070800" class="cur">营口</li><li ka="hot-city-101070900" data-val="101070900" class="cur">阜新</li><li ka="hot-city-101071000" data-val="101071000" class="cur">辽阳</li><li ka="hot-city-101071100" data-val="101071100" class="cur">铁岭</li><li ka="hot-city-101071200" data-val="101071200" class="cur">朝阳</li><li ka="hot-city-101071300" data-val="101071300" class="cur">盘锦</li><li ka="hot-city-101071400" data-val="101071400" class="cur">葫芦岛</li></ul><ul><li ka="hot-city-101080100" data-val="101080100" class="cur">呼和浩特</li><li ka="hot-city-101080200" data-val="101080200" class="cur">包头</li><li ka="hot-city-101080300" data-val="101080300" class="cur">乌海</li><li ka="hot-city-101080400" data-val="101080400" class="cur">通辽</li><li ka="hot-city-101080500" data-val="101080500" class="cur">赤峰</li><li ka="hot-city-101080600" data-val="101080600" class="cur">鄂尔多斯</li><li ka="hot-city-101080700" data-val="101080700" class="cur">呼伦贝尔</li><li ka="hot-city-101080800" data-val="101080800" class="cur">巴彦淖尔</li><li ka="hot-city-101080900" data-val="101080900" class="cur">乌兰察布</li><li ka="hot-city-101081000" data-val="101081000" class="cur">锡林郭勒</li><li ka="hot-city-101081100" data-val="101081100" class="cur">兴安盟</li><li ka="hot-city-101081200" data-val="101081200" class="cur">阿拉善</li></ul><ul><li ka="hot-city-101090100" data-val="101090100" class="cur">石家庄</li><li ka="hot-city-101090200" data-val="101090200" class="cur">保定</li><li ka="hot-city-101090300" data-val="101090300" class="cur">张家口</li><li ka="hot-city-101090400" data-val="101090400" class="cur">承德</li><li ka="hot-city-101090500" data-val="101090500" class="cur">唐山</li><li ka="hot-city-101090600" data-val="101090600" class="cur">廊坊</li><li ka="hot-city-101090700" data-val="101090700" class="cur">沧州</li><li ka="hot-city-101090800" data-val="101090800" class="cur">衡水</li><li ka="hot-city-101090900" data-val="101090900" class="cur">邢台</li><li ka="hot-city-101091000" data-val="101091000" class="cur">邯郸</li><li ka="hot-city-101091100" data-val="101091100" class="cur">秦皇岛</li></ul><ul><li ka="hot-city-101100100" data-val="101100100" class="cur">太原</li><li ka="hot-city-101100200" data-val="101100200" class="cur">大同</li><li ka="hot-city-101100300" data-val="101100300" class="cur">阳泉</li><li ka="hot-city-101100400" data-val="101100400" class="cur">晋中</li><li ka="hot-city-101100500" data-val="101100500" class="cur">长治</li><li ka="hot-city-101100600" data-val="101100600" class="cur">晋城</li><li ka="hot-city-101100700" data-val="101100700" class="cur">临汾</li><li ka="hot-city-101100800" data-val="101100800" class="cur">运城</li><li ka="hot-city-101100900" data-val="101100900" class="cur">朔州</li><li ka="hot-city-101101000" data-val="101101000" class="cur">忻州</li><li ka="hot-city-101101100" data-val="101101100" class="cur">吕梁</li></ul><ul><li ka="hot-city-101110100" data-val="101110100" class="cur">西安</li><li ka="hot-city-101110200" data-val="101110200" class="cur">咸阳</li><li ka="hot-city-101110300" data-val="101110300" class="cur">延安</li><li ka="hot-city-101110400" data-val="101110400" class="cur">榆林</li><li ka="hot-city-101110500" data-val="101110500" class="cur">渭南</li><li ka="hot-city-101110600" data-val="101110600" class="cur">商洛</li><li ka="hot-city-101110700" data-val="101110700" class="cur">安康</li><li ka="hot-city-101110800" data-val="101110800" class="cur">汉中</li><li ka="hot-city-101110900" data-val="101110900" class="cur">宝鸡</li><li ka="hot-city-101111000" data-val="101111000" class="cur">铜川</li></ul><ul><li ka="hot-city-101120100" data-val="101120100" class="cur">济南</li><li ka="hot-city-101120200" data-val="101120200" class="cur">青岛</li><li ka="hot-city-101120300" data-val="101120300" class="cur">淄博</li><li ka="hot-city-101120400" data-val="101120400" class="cur">德州</li><li ka="hot-city-101120500" data-val="101120500" class="cur">烟台</li><li ka="hot-city-101120600" data-val="101120600" class="cur">潍坊</li><li ka="hot-city-101120700" data-val="101120700" class="cur">济宁</li><li ka="hot-city-101120800" data-val="101120800" class="cur">泰安</li><li ka="hot-city-101120900" data-val="101120900" class="cur">临沂</li><li ka="hot-city-101121000" data-val="101121000" class="cur">菏泽</li><li ka="hot-city-101121100" data-val="101121100" class="cur">滨州</li><li ka="hot-city-101121200" data-val="101121200" class="cur">东营</li><li ka="hot-city-101121300" data-val="101121300" class="cur">威海</li><li ka="hot-city-101121400" data-val="101121400" class="cur">枣庄</li><li ka="hot-city-101121500" data-val="101121500" class="cur">日照</li><li ka="hot-city-101121700" data-val="101121700" class="cur">聊城</li></ul><ul><li ka="hot-city-101130100" data-val="101130100" class="cur">乌鲁木齐</li><li ka="hot-city-101130200" data-val="101130200" class="cur">克拉玛依</li><li ka="hot-city-101130300" data-val="101130300" class="cur">昌吉</li><li ka="hot-city-101130400" data-val="101130400" class="cur">巴音郭楞</li><li ka="hot-city-101130500" data-val="101130500" class="cur">博尔塔拉</li><li ka="hot-city-101130600" data-val="101130600" class="cur">伊犁</li><li ka="hot-city-101130800" data-val="101130800" class="cur">吐鲁番</li><li ka="hot-city-101130900" data-val="101130900" class="cur">哈密</li><li ka="hot-city-101131000" data-val="101131000" class="cur">阿克苏</li><li ka="hot-city-101131100" data-val="101131100" class="cur">克孜勒苏柯尔克孜</li><li ka="hot-city-101131200" data-val="101131200" class="cur">喀什</li><li ka="hot-city-101131300" data-val="101131300" class="cur">和田</li><li ka="hot-city-101131400" data-val="101131400" class="cur">塔城</li><li ka="hot-city-101131500" data-val="101131500" class="cur">阿勒泰</li><li ka="hot-city-101131600" data-val="101131600" class="cur">石河子</li><li ka="hot-city-101131700" data-val="101131700" class="cur">阿拉尔</li><li ka="hot-city-101131800" data-val="101131800" class="cur">图木舒克</li><li ka="hot-city-101131900" data-val="101131900" class="cur">五家渠</li><li ka="hot-city-101132000" data-val="101132000" class="cur">铁门关</li><li ka="hot-city-101132100" data-val="101132100" class="cur">北屯市</li><li ka="hot-city-101132200" data-val="101132200" class="cur">可克达拉市</li><li ka="hot-city-101132300" data-val="101132300" class="cur">昆玉市</li><li ka="hot-city-101132400" data-val="101132400" class="cur">双河市</li></ul><ul><li ka="hot-city-101150100" data-val="101150100" class="cur">西宁</li><li ka="hot-city-101150200" data-val="101150200" class="cur">海东</li><li ka="hot-city-101150300" data-val="101150300" class="cur">海北</li><li ka="hot-city-101150400" data-val="101150400" class="cur">黄南</li><li ka="hot-city-101150500" data-val="101150500" class="cur">海南</li><li ka="hot-city-101150600" data-val="101150600" class="cur">果洛</li><li ka="hot-city-101150700" data-val="101150700" class="cur">玉树</li><li ka="hot-city-101150800" data-val="101150800" class="cur">海西</li></ul><ul><li ka="hot-city-101160100" data-val="101160100" class="cur">兰州</li><li ka="hot-city-101160200" data-val="101160200" class="cur">定西</li><li ka="hot-city-101160300" data-val="101160300" class="cur">平凉</li><li ka="hot-city-101160400" data-val="101160400" class="cur">庆阳</li><li ka="hot-city-101160500" data-val="101160500" class="cur">武威</li><li ka="hot-city-101160600" data-val="101160600" class="cur">金昌</li><li ka="hot-city-101160700" data-val="101160700" class="cur">张掖</li><li ka="hot-city-101160800" data-val="101160800" class="cur">酒泉</li><li ka="hot-city-101160900" data-val="101160900" class="cur">天水</li><li ka="hot-city-101161000" data-val="101161000" class="cur">白银</li><li ka="hot-city-101161100" data-val="101161100" class="cur">陇南</li><li ka="hot-city-101161200" data-val="101161200" class="cur">嘉峪关</li><li ka="hot-city-101161300" data-val="101161300" class="cur">临夏</li><li ka="hot-city-101161400" data-val="101161400" class="cur">甘南</li></ul><ul><li ka="hot-city-101170100" data-val="101170100" class="cur">银川</li><li ka="hot-city-101170200" data-val="101170200" class="cur">石嘴山</li><li ka="hot-city-101170300" data-val="101170300" class="cur">吴忠</li><li ka="hot-city-101170400" data-val="101170400" class="cur">固原</li><li ka="hot-city-101170500" data-val="101170500" class="cur">中卫</li></ul><ul><li ka="hot-city-101180100" data-val="101180100" class="cur">郑州</li><li ka="hot-city-101180200" data-val="101180200" class="cur">安阳</li><li ka="hot-city-101180300" data-val="101180300" class="cur">新乡</li><li ka="hot-city-101180400" data-val="101180400" class="cur">许昌</li><li ka="hot-city-101180500" data-val="101180500" class="cur">平顶山</li><li ka="hot-city-101180600" data-val="101180600" class="cur">信阳</li><li ka="hot-city-101180700" data-val="101180700" class="cur">南阳</li><li ka="hot-city-101180800" data-val="101180800" class="cur">开封</li><li ka="hot-city-101180900" data-val="101180900" class="cur">洛阳</li><li ka="hot-city-101181000" data-val="101181000" class="cur">商丘</li><li ka="hot-city-101181100" data-val="101181100" class="cur">焦作</li><li ka="hot-city-101181200" data-val="101181200" class="cur">鹤壁</li><li ka="hot-city-101181300" data-val="101181300" class="cur">濮阳</li><li ka="hot-city-101181400" data-val="101181400" class="cur">周口</li><li ka="hot-city-101181500" data-val="101181500" class="cur">漯河</li><li ka="hot-city-101181600" data-val="101181600" class="cur">驻马店</li><li ka="hot-city-101181700" data-val="101181700" class="cur">三门峡</li><li ka="hot-city-101181800" data-val="101181800" class="cur">济源</li></ul><ul><li ka="hot-city-101190100" data-val="101190100" class="cur">南京</li><li ka="hot-city-101190200" data-val="101190200" class="cur">无锡</li><li ka="hot-city-101190300" data-val="101190300" class="cur">镇江</li><li ka="hot-city-101190400" data-val="101190400" class="cur">苏州</li><li ka="hot-city-101190500" data-val="101190500" class="cur">南通</li><li ka="hot-city-101190600" data-val="101190600" class="cur">扬州</li><li ka="hot-city-101190700" data-val="101190700" class="cur">盐城</li><li ka="hot-city-101190800" data-val="101190800" class="cur">徐州</li><li ka="hot-city-101190900" data-val="101190900" class="cur">淮安</li><li ka="hot-city-101191000" data-val="101191000" class="cur">连云港</li><li ka="hot-city-101191100" data-val="101191100" class="cur">常州</li><li ka="hot-city-101191200" data-val="101191200" class="cur">泰州</li><li ka="hot-city-101191300" data-val="101191300" class="cur">宿迁</li></ul><ul><li ka="hot-city-101200100" data-val="101200100" class="cur">武汉</li><li ka="hot-city-101200200" data-val="101200200" class="cur">襄阳</li><li ka="hot-city-101200300" data-val="101200300" class="cur">鄂州</li><li ka="hot-city-101200400" data-val="101200400" class="cur">孝感</li><li ka="hot-city-101200500" data-val="101200500" class="cur">黄冈</li><li ka="hot-city-101200600" data-val="101200600" class="cur">黄石</li><li ka="hot-city-101200700" data-val="101200700" class="cur">咸宁</li><li ka="hot-city-101200800" data-val="101200800" class="cur">荆州</li><li ka="hot-city-101200900" data-val="101200900" class="cur">宜昌</li><li ka="hot-city-101201000" data-val="101201000" class="cur">十堰</li><li ka="hot-city-101201100" data-val="101201100" class="cur">随州</li><li ka="hot-city-101201200" data-val="101201200" class="cur">荆门</li><li ka="hot-city-101201300" data-val="101201300" class="cur">恩施</li><li ka="hot-city-101201400" data-val="101201400" class="cur">仙桃</li><li ka="hot-city-101201500" data-val="101201500" class="cur">潜江</li><li ka="hot-city-101201600" data-val="101201600" class="cur">天门</li><li ka="hot-city-101201700" data-val="101201700" class="cur">神农架</li></ul><ul><li ka="hot-city-101210100" data-val="101210100" class="cur">杭州</li><li ka="hot-city-101210200" data-val="101210200" class="cur">湖州</li><li ka="hot-city-101210300" data-val="101210300" class="cur">嘉兴</li><li ka="hot-city-101210400" data-val="101210400" class="cur">宁波</li><li ka="hot-city-101210500" data-val="101210500" class="cur">绍兴</li><li ka="hot-city-101210600" data-val="101210600" class="cur">台州</li><li ka="hot-city-101210700" data-val="101210700" class="cur">温州</li><li ka="hot-city-101210800" data-val="101210800" class="cur">丽水</li><li ka="hot-city-101210900" data-val="101210900" class="cur">金华</li><li ka="hot-city-101211000" data-val="101211000" class="cur">衢州</li><li ka="hot-city-101211100" data-val="101211100" class="cur">舟山</li></ul><ul><li ka="hot-city-101220100" data-val="101220100" class="cur">合肥</li><li ka="hot-city-101220200" data-val="101220200" class="cur">蚌埠</li><li ka="hot-city-101220300" data-val="101220300" class="cur">芜湖</li><li ka="hot-city-101220400" data-val="101220400" class="cur">淮南</li><li ka="hot-city-101220500" data-val="101220500" class="cur">马鞍山</li><li ka="hot-city-101220600" data-val="101220600" class="cur">安庆</li><li ka="hot-city-101220700" data-val="101220700" class="cur">宿州</li><li ka="hot-city-101220800" data-val="101220800" class="cur">阜阳</li><li ka="hot-city-101220900" data-val="101220900" class="cur">亳州</li><li ka="hot-city-101221000" data-val="101221000" class="cur">滁州</li><li ka="hot-city-101221100" data-val="101221100" class="cur">淮北</li><li ka="hot-city-101221200" data-val="101221200" class="cur">铜陵</li><li ka="hot-city-101221300" data-val="101221300" class="cur">宣城</li><li ka="hot-city-101221400" data-val="101221400" class="cur">六安</li><li ka="hot-city-101221500" data-val="101221500" class="cur">池州</li><li ka="hot-city-101221600" data-val="101221600" class="cur">黄山</li></ul><ul><li ka="hot-city-101230100" data-val="101230100" class="cur">福州</li><li ka="hot-city-101230200" data-val="101230200" class="cur">厦门</li><li ka="hot-city-101230300" data-val="101230300" class="cur">宁德</li><li ka="hot-city-101230400" data-val="101230400" class="cur">莆田</li><li ka="hot-city-101230500" data-val="101230500" class="cur">泉州</li><li ka="hot-city-101230600" data-val="101230600" class="cur">漳州</li><li ka="hot-city-101230700" data-val="101230700" class="cur">龙岩</li><li ka="hot-city-101230800" data-val="101230800" class="cur">三明</li><li ka="hot-city-101230900" data-val="101230900" class="cur">南平</li></ul><ul><li ka="hot-city-101240100" data-val="101240100" class="cur">南昌</li><li ka="hot-city-101240200" data-val="101240200" class="cur">九江</li><li ka="hot-city-101240300" data-val="101240300" class="cur">上饶</li><li ka="hot-city-101240400" data-val="101240400" class="cur">抚州</li><li ka="hot-city-101240500" data-val="101240500" class="cur">宜春</li><li ka="hot-city-101240600" data-val="101240600" class="cur">吉安</li><li ka="hot-city-101240700" data-val="101240700" class="cur">赣州</li><li ka="hot-city-101240800" data-val="101240800" class="cur">景德镇</li><li ka="hot-city-101240900" data-val="101240900" class="cur">萍乡</li><li ka="hot-city-101241000" data-val="101241000" class="cur">新余</li><li ka="hot-city-101241100" data-val="101241100" class="cur">鹰潭</li></ul><ul><li ka="hot-city-101250100" data-val="101250100" class="cur">长沙</li><li ka="hot-city-101250200" data-val="101250200" class="cur">湘潭</li><li ka="hot-city-101250300" data-val="101250300" class="cur">株洲</li><li ka="hot-city-101250400" data-val="101250400" class="cur">衡阳</li><li ka="hot-city-101250500" data-val="101250500" class="cur">郴州</li><li ka="hot-city-101250600" data-val="101250600" class="cur">常德</li><li ka="hot-city-101250700" data-val="101250700" class="cur">益阳</li><li ka="hot-city-101250800" data-val="101250800" class="cur">娄底</li><li ka="hot-city-101250900" data-val="101250900" class="cur">邵阳</li><li ka="hot-city-101251000" data-val="101251000" class="cur">岳阳</li><li ka="hot-city-101251100" data-val="101251100" class="cur">张家界</li><li ka="hot-city-101251200" data-val="101251200" class="cur">怀化</li><li ka="hot-city-101251300" data-val="101251300" class="cur">永州</li><li ka="hot-city-101251400" data-val="101251400" class="cur">湘西</li></ul><ul><li ka="hot-city-101260100" data-val="101260100" class="cur">贵阳</li><li ka="hot-city-101260200" data-val="101260200" class="cur">遵义</li><li ka="hot-city-101260300" data-val="101260300" class="cur">安顺</li><li ka="hot-city-101260400" data-val="101260400" class="cur">铜仁</li><li ka="hot-city-101260500" data-val="101260500" class="cur">毕节</li><li ka="hot-city-101260600" data-val="101260600" class="cur">六盘水</li><li ka="hot-city-101260700" data-val="101260700" class="cur">黔东南</li><li ka="hot-city-101260800" data-val="101260800" class="cur">黔南</li><li ka="hot-city-101260900" data-val="101260900" class="cur">黔西南</li></ul><ul><li ka="hot-city-101270100" data-val="101270100" class="cur">成都</li><li ka="hot-city-101270200" data-val="101270200" class="cur">攀枝花</li><li ka="hot-city-101270300" data-val="101270300" class="cur">自贡</li><li ka="hot-city-101270400" data-val="101270400" class="cur">绵阳</li><li ka="hot-city-101270500" data-val="101270500" class="cur">南充</li><li ka="hot-city-101270600" data-val="101270600" class="cur">达州</li><li ka="hot-city-101270700" data-val="101270700" class="cur">遂宁</li><li ka="hot-city-101270800" data-val="101270800" class="cur">广安</li><li ka="hot-city-101270900" data-val="101270900" class="cur">巴中</li><li ka="hot-city-101271000" data-val="101271000" class="cur">泸州</li><li ka="hot-city-101271100" data-val="101271100" class="cur">宜宾</li><li ka="hot-city-101271200" data-val="101271200" class="cur">内江</li><li ka="hot-city-101271300" data-val="101271300" class="cur">资阳</li><li ka="hot-city-101271400" data-val="101271400" class="cur">乐山</li><li ka="hot-city-101271500" data-val="101271500" class="cur">眉山</li><li ka="hot-city-101271600" data-val="101271600" class="cur">雅安</li><li ka="hot-city-101271700" data-val="101271700" class="cur">德阳</li><li ka="hot-city-101271800" data-val="101271800" class="cur">广元</li><li ka="hot-city-101271900" data-val="101271900" class="cur">阿坝</li><li ka="hot-city-101272000" data-val="101272000" class="cur">凉山</li><li ka="hot-city-101272100" data-val="101272100" class="cur">甘孜</li></ul><ul><li ka="hot-city-101280100" data-val="101280100" class="cur">广州</li><li ka="hot-city-101280200" data-val="101280200" class="cur">韶关</li><li ka="hot-city-101280300" data-val="101280300" class="cur">惠州</li><li ka="hot-city-101280400" data-val="101280400" class="cur">梅州</li><li ka="hot-city-101280500" data-val="101280500" class="cur">汕头</li><li ka="hot-city-101280600" data-val="101280600" class="cur">深圳</li><li ka="hot-city-101280700" data-val="101280700" class="cur">珠海</li><li ka="hot-city-101280800" data-val="101280800" class="cur">佛山</li><li ka="hot-city-101280900" data-val="101280900" class="cur">肇庆</li><li ka="hot-city-101281000" data-val="101281000" class="cur">湛江</li><li ka="hot-city-101281100" data-val="101281100" class="cur">江门</li><li ka="hot-city-101281200" data-val="101281200" class="cur">河源</li><li ka="hot-city-101281300" data-val="101281300" class="cur">清远</li><li ka="hot-city-101281400" data-val="101281400" class="cur">云浮</li><li ka="hot-city-101281500" data-val="101281500" class="cur">潮州</li><li ka="hot-city-101281600" data-val="101281600" class="cur">东莞</li><li ka="hot-city-101281700" data-val="101281700" class="cur">中山</li><li ka="hot-city-101281800" data-val="101281800" class="cur">阳江</li><li ka="hot-city-101281900" data-val="101281900" class="cur">揭阳</li><li ka="hot-city-101282000" data-val="101282000" class="cur">茂名</li><li ka="hot-city-101282100" data-val="101282100" class="cur">汕尾</li><li ka="hot-city-101282200" data-val="101282200" class="cur">东沙群岛</li></ul><ul><li ka="hot-city-101290100" data-val="101290100" class="cur">昆明</li><li ka="hot-city-101290200" data-val="101290200" class="cur">曲靖</li><li ka="hot-city-101290300" data-val="101290300" class="cur">保山</li><li ka="hot-city-101290400" data-val="101290400" class="cur">玉溪</li><li ka="hot-city-101290500" data-val="101290500" class="cur">普洱</li><li ka="hot-city-101290700" data-val="101290700" class="cur">昭通</li><li ka="hot-city-101290800" data-val="101290800" class="cur">临沧</li><li ka="hot-city-101290900" data-val="101290900" class="cur">丽江</li><li ka="hot-city-101291000" data-val="101291000" class="cur">西双版纳</li><li ka="hot-city-101291100" data-val="101291100" class="cur">文山</li><li ka="hot-city-101291200" data-val="101291200" class="cur">红河</li><li ka="hot-city-101291300" data-val="101291300" class="cur">德宏</li><li ka="hot-city-101291400" data-val="101291400" class="cur">怒江</li><li ka="hot-city-101291500" data-val="101291500" class="cur">迪庆</li><li ka="hot-city-101291600" data-val="101291600" class="cur">大理</li><li ka="hot-city-101291700" data-val="101291700" class="cur">楚雄</li></ul><ul><li ka="hot-city-101300100" data-val="101300100" class="cur">南宁</li><li ka="hot-city-101300200" data-val="101300200" class="cur">崇左</li><li ka="hot-city-101300300" data-val="101300300" class="cur">柳州</li><li ka="hot-city-101300400" data-val="101300400" class="cur">来宾</li><li ka="hot-city-101300500" data-val="101300500" class="cur">桂林</li><li ka="hot-city-101300600" data-val="101300600" class="cur">梧州</li><li ka="hot-city-101300700" data-val="101300700" class="cur">贺州</li><li ka="hot-city-101300800" data-val="101300800" class="cur">贵港</li><li ka="hot-city-101300900" data-val="101300900" class="cur">玉林</li><li ka="hot-city-101301000" data-val="101301000" class="cur">百色</li><li ka="hot-city-101301100" data-val="101301100" class="cur">钦州</li><li ka="hot-city-101301200" data-val="101301200" class="cur">河池</li><li ka="hot-city-101301300" data-val="101301300" class="cur">北海</li><li ka="hot-city-101301400" data-val="101301400" class="cur">防城港</li></ul><ul><li ka="hot-city-101310100" data-val="101310100" class="cur">海口</li><li ka="hot-city-101310200" data-val="101310200" class="cur">三亚</li><li ka="hot-city-101310300" data-val="101310300" class="cur">三沙</li><li ka="hot-city-101310400" data-val="101310400" class="cur">儋州</li><li ka="hot-city-101310500" data-val="101310500" class="cur">五指山</li><li ka="hot-city-101310600" data-val="101310600" class="cur">琼海</li><li ka="hot-city-101310700" data-val="101310700" class="cur">文昌</li><li ka="hot-city-101310800" data-val="101310800" class="cur">万宁</li><li ka="hot-city-101310900" data-val="101310900" class="cur">东方</li><li ka="hot-city-101311000" data-val="101311000" class="cur">定安</li><li ka="hot-city-101311100" data-val="101311100" class="cur">屯昌</li><li ka="hot-city-101311200" data-val="101311200" class="cur">澄迈</li><li ka="hot-city-101311300" data-val="101311300" class="cur">临高</li><li ka="hot-city-101311400" data-val="101311400" class="cur">白沙</li><li ka="hot-city-101311500" data-val="101311500" class="cur">昌江</li><li ka="hot-city-101311600" data-val="101311600" class="cur">乐东</li><li ka="hot-city-101311700" data-val="101311700" class="cur">陵水</li><li ka="hot-city-101311800" data-val="101311800" class="cur">保亭</li><li ka="hot-city-101311900" data-val="101311900" class="cur">琼中</li></ul><ul><li ka="hot-city-101341100" data-val="101341100" class="cur">台湾</li></ul><ul><li ka="hot-city-101140100" data-val="101140100" class="cur">拉萨</li><li ka="hot-city-101140200" data-val="101140200" class="cur">日喀则</li><li ka="hot-city-101140300" data-val="101140300" class="cur">昌都</li><li ka="hot-city-101140400" data-val="101140400" class="cur">林芝</li><li ka="hot-city-101140500" data-val="101140500" class="cur">山南</li><li ka="hot-city-101140600" data-val="101140600" class="cur">那曲</li><li ka="hot-city-101140700" data-val="101140700" class="cur">阿里</li></ul><ul><li ka="hot-city-101320300" data-val="101320300" class="cur">香港</li></ul><ul><li ka="hot-city-101330100" data-val="101330100" class="cur">澳门</li></ul></div>
                                </div>
                            </form>
                        </div>
                    </div>
                         <div class="recommend-box">
                             <h3>推荐职位</h3>
                             <div class="slider-main" style="height: 390.78px;">
                                 <ul>
                                            <li class="cur">
                                        <a ka="job_recommend_1" href="/job_detail/dc133f9c9bc673251HV909m6F1M~.html" target="_blank" class="recommend-li">
                                            <div class="name">后端开发工程师（Go） <span class="red">10k-15k</span></div>
                                            <p>卓越工坊</p>
                                        </a>
                                        <a ka="job_recommend_2" href="/job_detail/59edf36b5ebb86981HZ-2926F1Q~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang后台开发工程师 <span class="red">18k-30k</span></div>
                                            <p>周同科技</p>
                                        </a>
                                        <a ka="job_recommend_3" href="/job_detail/c08c51ca1f5264671HV839q9F1A~.html" target="_blank" class="recommend-li">
                                            <div class="name">GO语言开发工程师 <span class="red">12k-20k</span></div>
                                            <p>盖娅互娱</p>
                                        </a>
                                        <a ka="job_recommend_4" href="/job_detail/2d0839578aa8ab4a1HR-09q5GVA~.html" target="_blank" class="recommend-li">
                                            <div class="name">golang开发工程师 <span class="red">20k-30k</span></div>
                                            <p>旷视科技</p>
                                        </a>
                                        <a ka="job_recommend_5" href="/job_detail/38675085d5f06cd11nZz3Ni7GVQ~.html" target="_blank" class="recommend-li">
                                            <div class="name">Lead Golang/Go开发工程师 <span class="red">35k-45k</span></div>
                                            <p>FreeWheel</p>
                                        </a>
                                        <a ka="job_recommend_6" href="/job_detail/af65e583a7e835b11XJ50tW9FVU~.html" target="_blank" class="recommend-li">
                                            <div class="name">go语言高级开发工程师 <span class="red">20k-30k</span></div>
                                            <p>摩点网</p>
                                        </a>
                                        <a ka="job_recommend_7" href="/job_detail/7c7532823fe29bf91Xx62dS9E1M~.html" target="_blank" class="recommend-li">
                                            <div class="name">go语言高级开发工程师 <span class="red">20k-40k</span></div>
                                            <p>她拍</p>
                                        </a>
                                        <a ka="job_recommend_8" href="/job_detail/85672d41f9ff174a1HZ-2Nm9GVI~.html" target="_blank" class="recommend-li">
                                            <div class="name">golang开发工程师 <span class="red">20k-40k</span></div>
                                            <p>今日头条</p>
                                        </a>
                                        <a ka="job_recommend_9" href="/job_detail/383f7951238edd051HZ72tm7GVU~.html" target="_blank" class="recommend-li">
                                            <div class="name">GO语言开发工程师 <span class="red">20k-40k</span></div>
                                            <p>伴鱼</p>
                                        </a>
                                        <a ka="job_recommend_10" href="/job_detail/5a949257ae9d7fde1nN62Nu6GFM~.html" target="_blank" class="recommend-li">
                                            <div class="name">Go开发工程师 <span class="red">25k-35k</span></div>
                                            <p>映客直播</p>
                                        </a>
                                            </li>
                                            <li class="">
                                        <a ka="job_recommend_11" href="/job_detail/f55f21e9dcbb4a201HZ62928GVs~.html" target="_blank" class="recommend-li">
                                            <div class="name">GO开发工程师 <span class="red">12k-18k</span></div>
                                            <p>包大师</p>
                                        </a>
                                        <a ka="job_recommend_12" href="/job_detail/dbea53ab01701fc01Xx72dy9Flo~.html" target="_blank" class="recommend-li">
                                            <div class="name">go语言开发工程师 <span class="red">15k-30k</span></div>
                                            <p>清檬养老</p>
                                        </a>
                                        <a ka="job_recommend_13" href="/job_detail/dd9256cb3e92acd01HR9096-FVM~.html" target="_blank" class="recommend-li">
                                            <div class="name">golang开发工程师 <span class="red">20k-40k</span></div>
                                            <p>格灵深瞳</p>
                                        </a>
                                        <a ka="job_recommend_14" href="/job_detail/84d58a83687c3bc31XV729i1GFU~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">20k-40k</span></div>
                                            <p>世纪好未来教育</p>
                                        </a>
                                        <a ka="job_recommend_15" href="/job_detail/86f6cdd71d21a0c41XZ62Nq-GVo~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">25k-50k</span></div>
                                            <p>宜信公司</p>
                                        </a>
                                        <a ka="job_recommend_16" href="/job_detail/c97f8a9d45ec516d1XNy0tm1GVE~.html" target="_blank" class="recommend-li">
                                            <div class="name">Go开发工程师 <span class="red">20k-40k</span></div>
                                            <p>小年糕</p>
                                        </a>
                                        <a ka="job_recommend_17" href="/job_detail/ac21d9566f0688301HRy3ti5FlE~.html" target="_blank" class="recommend-li">
                                            <div class="name">服务器端开发工程师（python/go） <span class="red">30k-60k</span></div>
                                            <p>今日头条</p>
                                        </a>
                                        <a ka="job_recommend_18" href="/job_detail/57b7dc2220d000661XB63N21ElE~.html" target="_blank" class="recommend-li">
                                            <div class="name">go开发工程师 <span class="red">25k-35k</span></div>
                                            <p>今日头条</p>
                                        </a>
                                        <a ka="job_recommend_19" href="/job_detail/1817d0c644bd1f611nN_0tW0E1Q~.html" target="_blank" class="recommend-li">
                                            <div class="name">C/C++ Java Golang开发工程师 <span class="red">15k-25k</span></div>
                                            <p>小爱智能科技</p>
                                        </a>
                                        <a ka="job_recommend_20" href="/job_detail/6c2a82ee6759186f1XB83924EVQ~.html" target="_blank" class="recommend-li">
                                            <div class="name">Go高级开发工程师 <span class="red">20k-40k</span></div>
                                            <p>波场</p>
                                        </a>
                                            </li>
                                            <li class="">
                                        <a ka="job_recommend_21" href="/job_detail/5f21612a08468ae61XN82965E1A~.html" target="_blank" class="recommend-li">
                                            <div class="name">golang开发工程师 <span class="red">13k-26k</span></div>
                                            <p>adhub</p>
                                        </a>
                                        <a ka="job_recommend_22" href="/job_detail/4bf9b1bedaec74141HV42d65FVc~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师（北京） <span class="red">8k-13k</span></div>
                                            <p>富麦科技</p>
                                        </a>
                                        <a ka="job_recommend_23" href="/job_detail/e0f050fd0aaab1ab1nNy3tq-EFE~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">15k-30k</span></div>
                                            <p>小年糕</p>
                                        </a>
                                        <a ka="job_recommend_24" href="/job_detail/1f6cb11c253c5b0c1XV93d67FFM~.html" target="_blank" class="recommend-li">
                                            <div class="name">go开发工程师 <span class="red">25k-50k</span></div>
                                            <p>瓜子二手车直卖网</p>
                                        </a>
                                        <a ka="job_recommend_25" href="/job_detail/75e1458e6cd2daf91HZ43Nm5GFc~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang后台开发工程师 <span class="red">15k-25k</span></div>
                                            <p>阿尔山区块链联盟</p>
                                        </a>
                                        <a ka="job_recommend_26" href="/job_detail/12f82ac2fb4842921HRz3ti-FVM~.html" target="_blank" class="recommend-li">
                                            <div class="name">golang开发工程师 <span class="red">15k-25k</span></div>
                                            <p>格灵深瞳</p>
                                        </a>
                                        <a ka="job_recommend_27" href="/job_detail/65adec2f254b6fbf1XJ939S0E1M~.html" target="_blank" class="recommend-li">
                                            <div class="name">服务端开发工程师（C++/golang） <span class="red">15k-30k</span></div>
                                            <p>360</p>
                                        </a>
                                        <a ka="job_recommend_28" href="/job_detail/10b21b5166dd225a1XF939q4FFA~.html" target="_blank" class="recommend-li">
                                            <div class="name">GO/Python 开发工程师 <span class="red">22k-35k</span></div>
                                            <p>一点资讯</p>
                                        </a>
                                        <a ka="job_recommend_29" href="/job_detail/20c6d316fbe4d1a41Xx42Nq0FFM~.html" target="_blank" class="recommend-li">
                                            <div class="name">Go语言开发工程师 <span class="red">8k-12k</span></div>
                                            <p>高弘科技</p>
                                        </a>
                                        <a ka="job_recommend_30" href="/job_detail/57dfec9985bcca7e1nx42dW8GFU~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">20k-40k</span></div>
                                            <p>轻松筹</p>
                                        </a>
                                            </li>
                                            <li class="">
                                        <a ka="job_recommend_31" href="/job_detail/a89abe2ddeda923a1nJ-3ty6EFA~.html" target="_blank" class="recommend-li">
                                            <div class="name">Go语言开发工程师 <span class="red">30k-40k</span></div>
                                            <p>云联万维</p>
                                        </a>
                                        <a ka="job_recommend_32" href="/job_detail/e257981f91672b0f1Hd429S5FlM~.html" target="_blank" class="recommend-li">
                                            <div class="name">后台开发工程师--Golang <span class="red">30k-60k</span></div>
                                            <p>今日头条</p>
                                        </a>
                                        <a ka="job_recommend_33" href="/job_detail/4281fcbff852a9a31HV-0ti9FVc~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang/后台开发工程师 <span class="red">15k-20k</span></div>
                                            <p>有鱼互娱</p>
                                        </a>
                                        <a ka="job_recommend_34" href="/job_detail/6096dab3cdf3c2d21n152tW8EFE~.html" target="_blank" class="recommend-li">
                                            <div class="name">GO开发工程师（IM业务） <span class="red">20k-35k</span></div>
                                            <p>巨洲云科技</p>
                                        </a>
                                        <a ka="job_recommend_35" href="/job_detail/ff0e2a76bc5907761Xd739-0F1s~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">20k-40k</span></div>
                                            <p>和风畅想</p>
                                        </a>
                                        <a ka="job_recommend_36" href="/job_detail/d1b1ff450a65cd271X162d69FlY~.html" target="_blank" class="recommend-li">
                                            <div class="name">高级服务器开发工程师（c++，golang） <span class="red">20k-35k</span></div>
                                            <p>掌通未来科技</p>
                                        </a>
                                        <a ka="job_recommend_37" href="/job_detail/456b38999991d7ca1HR729i_ElQ~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang开发工程师 <span class="red">30k-60k</span></div>
                                            <p>映客直播</p>
                                        </a>
                                        <a ka="job_recommend_38" href="/job_detail/e5a8a3340fb5c6e81HR-2N-_GFA~.html" target="_blank" class="recommend-li">
                                            <div class="name">后端开发工程师（Go/Golang） <span class="red">20k-35k</span></div>
                                            <p>医疗大数据</p>
                                        </a>
                                        <a ka="job_recommend_39" href="/job_detail/6f7defa6a19d78b11HR_29y_GVo~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang后端开发工程师 <span class="red">25k-40k</span></div>
                                            <p>映客直播</p>
                                        </a>
                                        <a ka="job_recommend_40" href="/job_detail/5d5372a0532135941X1y2N-0GVI~.html" target="_blank" class="recommend-li">
                                            <div class="name">Golang游戏开发工程师 <span class="red">25k-35k</span></div>
                                            <p>掌云网络科技</p>
                                        </a>
                                            </li>
                                 </ul>
                             <a href="javascript:;" class="btn-direction btn-prev prev"></a><a href="javascript:;" class="btn-direction btn-next next"></a><div class="slider-dot"><i data-index="1" class="cur"></i><i data-index="2"></i><i data-index="3"></i><i data-index="4"></i></div></div>
                         </div>
                    <!-- SEO  -->
<div class="links" style="height: 70px;">
        <dl class="links-item">
        <dt>推荐公司：</dt>
        <dd>
                        <a href="/gongsi/668c1d61d814564233B409s~.html" target="_blank" ka="seo-brand-1">熊小米</a>
                        <a href="/gongsi/5c2e5cac9caafc7633B4094~.html" target="_blank" ka="seo-brand-2">杭州飞牛科技</a>
                        <a href="/gongsi/93331487eb8903fc33B43Nw~.html" target="_blank" ka="seo-brand-3">时代</a>
                        <a href="/gongsi/927aef2198b71bca33B409U~.html" target="_blank" ka="seo-brand-4">友拓</a>
                        <a href="/gongsi/d2f2ab863b1393f833B409k~.html" target="_blank" ka="seo-brand-5">玖城在线</a>
                        <a href="/gongsi/61bf60a8ebf2d2331n1829W6.html" target="_blank" ka="seo-brand-6">北京固瑞恩</a>
                        <a href="/gongsi/859142d3a4c379721XR539S1.html" target="_blank" ka="seo-brand-7">超流利</a>
                        <a href="/gongsi/5bbe5e3e978a564c03V73ds~.html" target="_blank" ka="seo-brand-8">艾佳健身</a>
                        <a href="/gongsi/046bd7f1e3fcf52b1XR-3Nm4.html" target="_blank" ka="seo-brand-9">亿霖投资公司</a>
                        <a href="/gongsi/5dbb133ed8523e911Xd-2Ny7.html" target="_blank" ka="seo-brand-10">慧信</a>
                        <a href="/gongsi/2ab1349d0c019ab503R439S_.html" target="_blank" ka="seo-brand-11">杭州弘历软件</a>
                        <a href="/gongsi/1365fccd7ad4d1af1HNz09k~.html" target="_blank" ka="seo-brand-12">瞬联软件科技</a>
                        <a href="/gongsi/b9f613c220ece7161nd-3ti4GA~~.html" target="_blank" ka="seo-brand-13">华锐联创</a>
                        <a href="/gongsi/65b4086bd89fec771nd539S7.html" target="_blank" ka="seo-brand-14">久圣投资管理</a>
                        <a href="/gongsi/32c06995b056ebd903Fz3NU~.html" target="_blank" ka="seo-brand-15">华庭装饰</a>
                        <a href="/gongsi/3ba18cdf9650ebcd0nB-29q8.html" target="_blank" ka="seo-brand-16">鱼乐贝贝婴幼儿水育馆</a>
                        <a href="/gongsi/fc020e045d2cd8ba1XV539y4.html" target="_blank" ka="seo-brand-17">晟丰原科技</a>
                        <a href="/gongsi/06f32129e188d7931HFy2Ni4.html" target="_blank" ka="seo-brand-18">小区无忧</a>
                        <a href="/gongsi/ae3fd4d52085fa101nB73di_.html" target="_blank" ka="seo-brand-19">自由自宅</a>
                        <a href="/gongsi/3f9bf0259dc1e7b11nJ73dW_FA~~.html" target="_blank" ka="seo-brand-20">宅自由公寓</a>
                        <a href="/gongsi/beabfcfc549fc0d31XN_29y8.html" target="_blank" ka="seo-brand-21">网小二</a>
                        <a href="/gongsi/18d1e3555c1d09b11nxy39u-.html" target="_blank" ka="seo-brand-22">综佳贸易</a>
                        <a href="/gongsi/d5c3c0a62d58818603xz3t-1.html" target="_blank" ka="seo-brand-23">新力渠道部</a>
                        <a href="/gongsi/ac3130be2b9b77f91nN92967.html" target="_blank" ka="seo-brand-24">追尚网络科技</a>
                        <a href="/gongsi/73887ea24096fbe51nx50967.html" target="_blank" ka="seo-brand-25">长松咨询</a>
                        <a href="/gongsi/cd0ed90cce2c5b6d1nZ609y0GA~~.html" target="_blank" ka="seo-brand-26">国信假日酒店</a>
                        <a href="/gongsi/e0ad7f23108d40361HV409k~.html" target="_blank" ka="seo-brand-27">北京长松咨询</a>
                        <a href="/gongsi/ba7a79c5fc54d6651Xd839u9.html" target="_blank" ka="seo-brand-28">壹上悦农业</a>
                        <a href="/gongsi/f80691bdf3e429581nx72di4Fw~~.html" target="_blank" ka="seo-brand-29">鱼乐贝贝婴幼儿水族馆</a>
                        <a href="/gongsi/45884a62d18c9bcd0XVz3dy_.html" target="_blank" ka="seo-brand-30">辽宁善金</a>
        </dd>
    </dl>
    <dl class="links-item">
        <dt>城市招聘：</dt>
        <dd>
                     <a href="/c101071300/" target="_blank" ka="seo-city-1">盘锦招聘</a>
                     <a href="/c101120200/" target="_blank" ka="seo-city-2">青岛招聘</a>
                     <a href="/c101090200/" target="_blank" ka="seo-city-3">保定招聘</a>
                     <a href="/c101120800/" target="_blank" ka="seo-city-4">泰安招聘</a>
                     <a href="/c101051000/" target="_blank" ka="seo-city-5">鸡西招聘</a>
                     <a href="/c101070400/" target="_blank" ka="seo-city-6">抚顺招聘</a>
                     <a href="/c101210400/" target="_blank" ka="seo-city-7">宁波招聘</a>
                     <a href="/c101270200/" target="_blank" ka="seo-city-8">攀枝花招聘</a>
                     <a href="/c101311400/" target="_blank" ka="seo-city-9">白沙招聘</a>
                     <a href="/c101140700/" target="_blank" ka="seo-city-10">阿里招聘</a>
                     <a href="/c101291300/" target="_blank" ka="seo-city-11">德宏招聘</a>
                     <a href="/c101240400/" target="_blank" ka="seo-city-12">抚州招聘</a>
                     <a href="/c101060600/" target="_blank" ka="seo-city-13">辽源招聘</a>
                     <a href="/c101050800/" target="_blank" ka="seo-city-14">大庆招聘</a>
                     <a href="/c101140400/" target="_blank" ka="seo-city-15">林芝招聘</a>
                     <a href="/c101060900/" target="_blank" ka="seo-city-16">延边招聘</a>
                     <a href="/c101100400/" target="_blank" ka="seo-city-17">晋中招聘</a>
                     <a href="/c101170200/" target="_blank" ka="seo-city-18">石嘴山招聘</a>
                     <a href="/c101241100/" target="_blank" ka="seo-city-19">鹰潭招聘</a>
                     <a href="/c101230200/" target="_blank" ka="seo-city-20">厦门招聘</a>
        </dd>
    </dl>
    <dl class="links-item">
        <dt>热门职位：</dt>
        <dd>
                    <a href="/p101008/" target="_blank" ka="seo-position-1">无线射频工程师招聘</a>
                    <a href="/p100106/" target="_blank" ka="seo-position-2">C#招聘</a>
                    <a href="/p100401/" target="_blank" ka="seo-position-3">运维工程师招聘</a>
                    <a href="/p100702/" target="_blank" ka="seo-position-4">技术总监招聘</a>
                    <a href="/p100606/" target="_blank" ka="seo-position-5">实施工程师招聘</a>
                    <a href="/p100305/" target="_blank" ka="seo-position-6">测试开发招聘</a>
                    <a href="/p101016/" target="_blank" ka="seo-position-7">光通信工程师招聘</a>
                    <a href="/p101202/" target="_blank" ka="seo-position-8">售后工程师招聘</a>
                    <a href="/p100512/" target="_blank" ka="seo-position-9">数据架构师招聘</a>
                    <a href="/p101401/" target="_blank" ka="seo-position-10">电子工程师招聘</a>
                    <a href="/p100901/" target="_blank" ka="seo-position-11">web前端招聘</a>
                    <a href="/p100802/" target="_blank" ka="seo-position-12">嵌入式招聘</a>
                    <a href="/p101018/" target="_blank" ka="seo-position-13">光网络工程师招聘</a>
                    <a href="/p100405/" target="_blank" ka="seo-position-14">IT技术支持招聘</a>
                    <a href="/p100599/" target="_blank" ka="seo-position-15">数据招聘</a>
                    <a href="/p101003/" target="_blank" ka="seo-position-16">数据通信工程师招聘</a>
                    <a href="/p100816/" target="_blank" ka="seo-position-17">射频工程师招聘</a>
                    <a href="/p100109/" target="_blank" ka="seo-position-18">Python招聘</a>
                    <a href="/p100113/" target="_blank" ka="seo-position-19">Ruby招聘</a>
                    <a href="/p100813/" target="_blank" ka="seo-position-20">热传导招聘</a>
                    <a href="/p100203/" target="_blank" ka="seo-position-21">iOS招聘</a>
                    <a href="/p101399/" target="_blank" ka="seo-position-22">人工智能招聘</a>
                    <a href="/p100206/" target="_blank" ka="seo-position-23">Flash招聘</a>
                    <a href="/p100205/" target="_blank" ka="seo-position-24">移动web前端招聘</a>
                    <a href="/p101301/" target="_blank" ka="seo-position-25">机器学习招聘</a>
                    <a href="/p100108/" target="_blank" ka="seo-position-26">Hadoop招聘</a>
                    <a href="/p100799/" target="_blank" ka="seo-position-27">高端技术职位招聘</a>
                    <a href="/p100210/" target="_blank" ka="seo-position-28">COCOS2DX招聘</a>
                    <a href="/p100602/" target="_blank" ka="seo-position-29">项目主管招聘</a>
                    <a href="/p100306/" target="_blank" ka="seo-position-30">移动端测试招聘</a>
        </dd>
    </dl>
    <label><span>展开</span><i class="fz fz-slidedown"></i></label>
</div>                    <div class="pos-bread">
                        <a href="/" ka="job-breadcrumb-bottom1">首页</a><i class="icon-arrow-right"></i><a href="/c101010100/" ka="job-breadcrumb-bottom2">北京招聘</a>·<a href="/p100116/" ka="job-breadcrumb-bottom21">Golang招聘</a><i class="icon-arrow-right"></i><a href="/c101010100-p100116/" ka="job-breadcrumb-bottom3">北京Golang招聘</a><i class="icon-arrow-right"></i><a href="/job_detail/156cdef2df73d1391Xxy2ti0GFc~.html" ka="job-breadcrumb-bottom4">小米Golang招聘</a>
                    </div>
                </div>
            </div>
        </div>
    </div>
</div>`
	parseResult, err := ParsePosition([]byte(html))

	if err != nil {
		t.Errorf("parse error : %s", err.Error())
	}
	item := parseResult.Item
	if item.Positon == "" {
		t.Error("position is empty")
	}
	if item.Salary == "" {
		t.Error("salary is empty")
	}
	if item.City == "" {
		t.Error("city is empty")
	}
	if item.WorkYear == "" {
		t.Error("workYear is empty")
	}
	if item.Education == "" {
		t.Error("education is empty")
	}
	if item.CompanyName == "" {
		t.Error("companyName is empty")
	}
	if item.Industry == "" {
		t.Error("industry is empty")
	}
	if item.FinanceStage == "" {
		t.Error("financeStage is empty")
	}
	if item.CompanySize == "" {
		t.Error("companySize is empty")
	}
	if item.Detail == "" {
		t.Error("detail is empty")
	}
	if item.Location == "" {
		t.Error("location is empty")
	}
}
