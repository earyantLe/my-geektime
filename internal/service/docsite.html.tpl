<!DOCTYPE html>
<html lang="zh-CN">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <meta name="referrer" content="no-referrer" />
    <title>{{.Title}} - {{.SiteTitle}}</title>
    <style>
        :root {
            --ifm-color-primary: #1a73e8;
            --ifm-color-primary-dark: #155dbd;
            --ifm-color-primary-darker: #1458b3;
            --ifm-color-primary-darkest: #104993;
            --ifm-color-primary-light: #4c8df0;
            --ifm-color-primary-lighter: #5b96f1;
            --ifm-color-primary-lightest: #7cb0f5;
            
            --ifm-background-color: #ffffff;
            --ifm-background-surface-color: #f6f7f8;
            --ifm-font-color-base: #1c1e21;
            --ifm-font-color-secondary: #606770;
            --ifm-border-color: #ebedf0;
            --ifm-code-background: #f6f7f8;
            --ifm-blockquote-border-color: #1a73e8;
            
            --ifm-navbar-height: 60px;
            --ifm-sidebar-width: 300px;
            --ifm-container-max-width: 1140px;
            
            --ifm-font-family-base: -apple-system, BlinkMacSystemFont, "Segoe UI", Roboto, "Helvetica Neue", Arial, sans-serif;
            --ifm-font-family-monospace: SFMono-Regular, Menlo, Monaco, Consolas, monospace;
            
            --ifm-global-spacing: 1rem;
            --ifm-transition-fast: 200ms;
            --ifm-transition-slow: 400ms;
        }
        
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
        }
        
        html {
            scroll-behavior: smooth;
        }
        
        body {
            font-family: var(--ifm-font-family-base);
            line-height: 1.65;
            color: var(--ifm-font-color-base);
            background-color: var(--ifm-background-color);
            -webkit-font-smoothing: antialiased;
            -moz-osx-font-smoothing: grayscale;
        }
        
        /* 导航栏 */
        .navbar {
            position: fixed;
            top: 0;
            left: 0;
            right: 0;
            height: var(--ifm-navbar-height);
            background: rgba(255, 255, 255, 0.95);
            backdrop-filter: blur(10px);
            border-bottom: 1px solid var(--ifm-border-color);
            display: flex;
            align-items: center;
            padding: 0 1.5rem;
            z-index: 1000;
            transition: all var(--ifm-transition-fast) ease;
        }
        
        .navbar-brand {
            display: flex;
            align-items: center;
            gap: 0.75rem;
            text-decoration: none;
            color: var(--ifm-font-color-base);
            font-weight: 700;
            font-size: 1.25rem;
        }
        
        .navbar-brand:hover {
            color: var(--ifm-color-primary);
        }
        
        .navbar-logo {
            width: 32px;
            height: 32px;
            border-radius: 6px;
        }
        
        .menu-toggle {
            display: none;
            background: none;
            border: none;
            color: var(--ifm-font-color-base);
            font-size: 1.5rem;
            cursor: pointer;
            padding: 0.5rem;
            border-radius: 4px;
            transition: background var(--ifm-transition-fast) ease;
        }
        
        .menu-toggle:hover {
            background: var(--ifm-background-surface-color);
        }
        
        /* 主布局 */
        .main-wrapper {
            display: flex;
            margin-top: var(--ifm-navbar-height);
            min-height: calc(100vh - var(--ifm-navbar-height));
        }
        
        /* 侧边栏 */
        .sidebar {
            width: var(--ifm-sidebar-width);
            background: var(--ifm-background-color);
            border-right: 1px solid var(--ifm-border-color);
            position: fixed;
            top: var(--ifm-navbar-height);
            left: 0;
            bottom: 0;
            overflow-y: auto;
            padding: 1.5rem 0;
            transition: transform var(--ifm-transition-fast) ease;
            scrollbar-width: thin;
            scrollbar-color: var(--ifm-border-color) transparent;
        }
        
        .sidebar::-webkit-scrollbar {
            width: 6px;
        }
        
        .sidebar::-webkit-scrollbar-track {
            background: transparent;
        }
        
        .sidebar::-webkit-scrollbar-thumb {
            background: var(--ifm-border-color);
            border-radius: 3px;
        }
        
        .sidebar-section {
            margin-bottom: 1.5rem;
        }
        
        .sidebar-section-title {
            font-size: 0.75rem;
            font-weight: 600;
            color: var(--ifm-font-color-secondary);
            padding: 0.5rem 1.5rem;
            text-transform: uppercase;
            letter-spacing: 0.5px;
            margin-bottom: 0.5rem;
        }
        
        .nav-list {
            list-style: none;
            padding: 0;
        }
        
        .nav-item {
            margin: 0;
        }
        
        .nav-link {
            display: block;
            padding: 0.6rem 1.5rem;
            color: var(--ifm-font-color-base);
            text-decoration: none;
            font-size: 0.95rem;
            line-height: 1.5;
            border-left: 3px solid transparent;
            transition: all var(--ifm-transition-fast) ease;
        }
        
        .nav-link:hover {
            background: var(--ifm-background-surface-color);
            color: var(--ifm-color-primary);
            border-left-color: var(--ifm-color-primary-light);
        }
        
        .nav-link.active {
            background: linear-gradient(90deg, rgba(26, 115, 232, 0.08) 0%, transparent 100%);
            color: var(--ifm-color-primary);
            border-left-color: var(--ifm-color-primary);
            font-weight: 600;
        }
        
        /* 主内容区 */
        .main-content {
            margin-left: var(--ifm-sidebar-width);
            flex: 1;
            max-width: calc(100% - var(--ifm-sidebar-width));
            padding: 2.5rem 3rem;
        }
        
        .container {
            max-width: 900px;
            margin: 0 auto;
        }
        
        /* 文章头部 */
        .article-header {
            margin-bottom: 2.5rem;
            padding-bottom: 1.5rem;
            border-bottom: 2px solid var(--ifm-border-color);
        }
        
        .article-title {
            font-size: 2.5rem;
            font-weight: 800;
            color: var(--ifm-font-color-base);
            margin-bottom: 1rem;
            line-height: 1.3;
        }
        
        .article-meta {
            display: flex;
            align-items: center;
            gap: 1rem;
            color: var(--ifm-font-color-secondary);
            font-size: 0.9rem;
        }
        
        .badge {
            display: inline-block;
            padding: 0.25rem 0.75rem;
            background: var(--ifm-color-primary);
            color: white;
            border-radius: 12px;
            font-size: 0.8rem;
            font-weight: 600;
        }
        
        /* Markdown 内容样式 */
        .markdown-body {
            font-size: 1.05rem;
            line-height: 1.8;
            color: var(--ifm-font-color-base);
        }
        
        .markdown-body h1,
        .markdown-body h2,
        .markdown-body h3,
        .markdown-body h4,
        .markdown-body h5,
        .markdown-body h6 {
            margin-top: 2.5rem;
            margin-bottom: 1rem;
            font-weight: 700;
            line-height: 1.3;
            color: var(--ifm-font-color-base);
        }
        
        .markdown-body h1 { 
            font-size: 2rem; 
            border-bottom: 2px solid var(--ifm-border-color); 
            padding-bottom: 0.5rem;
            margin-top: 0;
        }
        
        .markdown-body h2 { 
            font-size: 1.65rem; 
            border-bottom: 1px solid var(--ifm-border-color); 
            padding-bottom: 0.4rem;
        }
        
        .markdown-body h3 { font-size: 1.35rem; }
        .markdown-body h4 { font-size: 1.15rem; }
        .markdown-body h5 { font-size: 1rem; }
        .markdown-body h6 { font-size: 0.9rem; color: var(--ifm-font-color-secondary); }
        
        .markdown-body p {
            margin-bottom: 1.25rem;
        }
        
        .markdown-body a {
            color: var(--ifm-color-primary);
            text-decoration: none;
            border-bottom: 1px solid transparent;
            transition: border-color var(--ifm-transition-fast) ease;
        }
        
        .markdown-body a:hover {
            border-bottom-color: var(--ifm-color-primary);
        }
        
        .markdown-body code {
            background: var(--ifm-code-background);
            padding: 0.2em 0.4em;
            border-radius: 3px;
            font-family: var(--ifm-font-family-monospace);
            font-size: 0.9em;
            color: #e83e8c;
        }
        
        .markdown-body pre {
            background: #282c34;
            color: #abb2bf;
            padding: 1.25rem;
            border-radius: 8px;
            overflow: auto;
            margin-bottom: 1.5rem;
            line-height: 1.6;
        }
        
        .markdown-body pre code {
            background: transparent;
            padding: 0;
            color: inherit;
            font-size: 0.9rem;
        }
        
        .markdown-body blockquote {
            border-left: 4px solid var(--ifm-blockquote-border-color);
            padding: 1rem 1.5rem;
            margin: 1.5rem 0;
            background: var(--ifm-background-surface-color);
            border-radius: 0 6px 6px 0;
            color: var(--ifm-font-color-secondary);
        }
        
        .markdown-body blockquote p:last-child {
            margin-bottom: 0;
        }
        
        .markdown-body ul,
        .markdown-body ol {
            margin-bottom: 1.25rem;
            padding-left: 2em;
        }
        
        .markdown-body li {
            margin: 0.5rem 0;
        }
        
        .markdown-body table {
            border-collapse: collapse;
            width: 100%;
            margin-bottom: 1.5rem;
            border-radius: 6px;
            overflow: hidden;
            box-shadow: 0 1px 3px rgba(0, 0, 0, 0.05);
        }
        
        .markdown-body th,
        .markdown-body td {
            border: 1px solid var(--ifm-border-color);
            padding: 0.75rem 1rem;
            text-align: left;
        }
        
        .markdown-body th {
            background: var(--ifm-background-surface-color);
            font-weight: 600;
            color: var(--ifm-font-color-base);
        }
        
        .markdown-body tr:nth-child(even) {
            background: var(--ifm-background-surface-color);
        }
        
        .markdown-body img {
            max-width: 100%;
            height: auto;
            border-radius: 8px;
            margin: 1.5rem 0;
            box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
        }
        
        /* 视频和音频播放器 */
        .media-player-container {
            position: relative;
            width: 100%;
            margin: 1.5rem 0;
            background: #000;
            border-radius: 8px;
            overflow: hidden;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
        }
        
        .media-player-inline {
            width: 100%;
            max-height: 500px;
        }
        
        .media-player-inline video,
        .media-player-inline audio {
            width: 100%;
            display: block;
        }
        
        /* 浮动播放器 */
        .media-player-float {
            position: fixed;
            bottom: 20px;
            right: 20px;
            width: 320px;
            background: #000;
            border-radius: 8px;
            box-shadow: 0 4px 20px rgba(0, 0, 0, 0.3);
            z-index: 9999;
            transition: all 0.3s ease;
            overflow: hidden;
        }
        
        .media-player-float video,
        .media-player-float audio {
            width: 100%;
            display: block;
        }
        
        .media-player-float .float-close-btn {
            position: absolute;
            top: 8px;
            right: 8px;
            width: 28px;
            height: 28px;
            background: rgba(0, 0, 0, 0.7);
            border: none;
            border-radius: 50%;
            color: white;
            font-size: 18px;
            cursor: pointer;
            display: flex;
            align-items: center;
            justify-content: center;
            transition: background 0.2s ease;
            z-index: 10000;
        }
        
        .media-player-float .float-close-btn:hover {
            background: rgba(0, 0, 0, 0.9);
        }
        
        .media-player-float.hidden {
            opacity: 0;
            pointer-events: none;
            transform: translateY(20px);
        }
        
        /* 原始播放器位置占位符 */
        .media-player-placeholder {
            width: 100%;
            min-height: 200px;
            background: #f0f0f0;
            border-radius: 8px;
            display: flex;
            align-items: center;
            justify-content: center;
            color: var(--ifm-font-color-secondary);
            font-size: 0.9rem;
        }
        
        /* 评论区 */
        .comments-section {
            margin-top: 4rem;
            padding-top: 2.5rem;
            border-top: 2px solid var(--ifm-border-color);
        }
        
        .comments-section h3 {
            font-size: 1.5rem;
            margin-bottom: 1.5rem;
            color: var(--ifm-font-color-base);
            display: flex;
            align-items: center;
            gap: 0.5rem;
        }
        
        .comment-list {
            list-style: none;
            padding: 0;
        }
        
        .comment-item {
            padding: 1.25rem;
            margin-bottom: 1rem;
            background: var(--ifm-background-surface-color);
            border-radius: 8px;
            border: 1px solid var(--ifm-border-color);
            transition: all var(--ifm-transition-fast) ease;
        }
        
        .comment-item:hover {
            box-shadow: 0 2px 8px rgba(0, 0, 0, 0.08);
            transform: translateY(-2px);
        }
        
        .comment-header {
            display: flex;
            align-items: center;
            margin-bottom: 0.75rem;
        }
        
        .comment-avatar {
            width: 40px;
            height: 40px;
            border-radius: 50%;
            margin-right: 0.75rem;
            border: 2px solid white;
            box-shadow: 0 1px 3px rgba(0, 0, 0, 0.1);
        }
        
        .comment-user {
            font-weight: 600;
            color: var(--ifm-font-color-base);
            font-size: 0.95rem;
        }
        
        .comment-stats {
            margin-left: auto;
            color: var(--ifm-font-color-secondary);
            font-size: 0.85rem;
            display: flex;
            gap: 1rem;
        }
        
        .comment-content {
            margin-top: 0.75rem;
            color: var(--ifm-font-color-base);
            line-height: 1.6;
        }
        
        .comment-date {
            margin-top: 0.5rem;
            color: var(--ifm-font-color-secondary);
            font-size: 0.8rem;
        }
        
        /* 分页导航 */
        .pagination {
            display: flex;
            justify-content: space-between;
            margin-top: 4rem;
            padding-top: 2rem;
            border-top: 1px solid var(--ifm-border-color);
            gap: 1rem;
        }
        
        .pagination a,
        .pagination span {
            padding: 0.75rem 1.75rem;
            background: var(--ifm-color-primary);
            color: #fff;
            text-decoration: none;
            border-radius: 50px;
            transition: all var(--ifm-transition-fast) ease;
            font-size: 0.95rem;
            font-weight: 500;
            display: inline-flex;
            align-items: center;
            gap: 0.5rem;
        }
        
        .pagination a:hover {
            background: var(--ifm-color-primary-dark);
            transform: translateY(-2px);
            box-shadow: 0 4px 12px rgba(26, 115, 232, 0.3);
        }
        
        .pagination .disabled {
            opacity: 0.5;
            pointer-events: none;
            background: var(--ifm-border-color);
            color: var(--ifm-font-color-secondary);
        }
        
        /* 响应式设计 */
        @media (max-width: 1024px) {
            .sidebar {
                width: 260px;
            }
            
            .main-content {
                margin-left: 260px;
                max-width: calc(100% - 260px);
                padding: 2rem;
            }
        }
        
        @media (max-width: 768px) {
            .menu-toggle {
                display: block;
            }
            
            .sidebar {
                transform: translateX(-100%);
                z-index: 999;
                box-shadow: 2px 0 12px rgba(0, 0, 0, 0.1);
                width: 280px;
            }
            
            .sidebar.open {
                transform: translateX(0);
            }
            
            .main-content {
                margin-left: 0;
                max-width: 100%;
                padding: 1.5rem;
            }
            
            .article-title {
                font-size: 2rem;
            }
            
            .markdown-body {
                font-size: 1rem;
            }
            
            .pagination {
                flex-direction: column;
            }
            
            .pagination a,
            .pagination span {
                width: 100%;
                justify-content: center;
            }
        }
        
        /* 动画效果 */
        @keyframes fadeIn {
            from {
                opacity: 0;
                transform: translateY(10px);
            }
            to {
                opacity: 1;
                transform: translateY(0);
            }
        }
        
        .main-content {
            animation: fadeIn 0.4s ease-out;
        }
        
        /* 打印样式 */
        @media print {
            .navbar,
            .sidebar,
            .pagination,
            .comments-section {
                display: none;
            }
            
            .main-content {
                margin-left: 0;
                max-width: 100%;
            }
        }
    </style>
    
    <!-- HLS.js for HLS video support -->
    <script src="https://cdn.jsdelivr.net/npm/hls.js@latest"></script>
</head>
<body>
    <!-- 顶部导航栏 -->
    <nav class="navbar">
        <button class="menu-toggle" onclick="toggleSidebar()" aria-label="Toggle menu">☰</button>
        <a href="index.html" class="navbar-brand">
            <span>{{.SiteTitle}}</span>
        </a>
    </nav>
    
    <div class="main-wrapper">
        <!-- 侧边栏导航 -->
        <aside class="sidebar" id="sidebar">
            <div class="sidebar-section">
                <div class="sidebar-section-title">📚 课程目录</div>
                <ul class="nav-list">
                    <li class="nav-item">
                        <a href="index.html" class="nav-link{{if .IsIndex}} active{{end}}">
                            🏠 首页概览
                        </a>
                    </li>
                    {{range .Articles}}
                    <li class="nav-item">
                        <a href="{{.Filename}}" class="nav-link{{if eq .Index $.CurrentIndex}} active{{end}}">
                            <span>{{.Index}}.</span> {{.Title}}
                        </a>
                    </li>
                    {{end}}
                </ul>
            </div>
        </aside>
        
        <!-- 主内容区 -->
        <main class="main-content">
            <div class="container">
                <header class="article-header">
                    <h1 class="article-title">{{.Title}}</h1>
                    {{if .ArticleIndex}}
                    <div class="article-meta">
                        <span class="badge">第 {{.ArticleIndex}} / {{.TotalArticles}} 篇</span>
                        <span>📖 阅读时间约 {{div (len .Content) 300}} 分钟</span>
                    </div>
                    {{end}}
                </header>
                
                <article class="markdown-body">
                    {{if .MediaPlayer}}
                    <div class="media-player-wrapper" style="margin-bottom: 2rem;">
                        {{.MediaPlayer}}
                    </div>
                    {{end}}
                    {{.Content}}
                </article>
                
                {{if .CommentsHTML}}
                <section class="comments-section">
                    <h3>💬 精选留言</h3>
                    <ul class="comment-list">
                        {{.CommentsHTML}}
                    </ul>
                </section>
                {{end}}
                
                <!-- 分页导航 -->
                {{if not .IsIndex}}
                <nav class="pagination">
                    {{if .PrevURL}}
                    <a href="{{.PrevURL}}">
                        ← 上一篇
                    </a>
                    {{else}}
                    <span class="disabled">← 上一篇</span>
                    {{end}}
                    
                    {{if .NextURL}}
                    <a href="{{.NextURL}}">
                        下一篇 →
                    </a>
                    {{else}}
                    <span class="disabled">下一篇 →</span>
                    {{end}}
                </nav>
                {{end}}
            </div>
        </main>
    </div>
    
    <script>
        // HLS 播放器初始化
        (function() {
            const videos = document.querySelectorAll('video');
            videos.forEach(function(video) {
                const source = video.querySelector('source');
                if (!source) return;
                
                const src = source.getAttribute('src');
                if (!src || !src.includes('.m3u8')) return;
                
                // 如果浏览器支持 HLS.js
                if (typeof Hls !== 'undefined' && Hls.isSupported()) {
                    const hls = new Hls({
                        enableWorker: true,
                        lowLatencyMode: true,
                        debug: false
                    });
                    
                    hls.loadSource(src);
                    hls.attachMedia(video);
                    
                    hls.on(Hls.Events.ERROR, function(event, data) {
                        console.error('HLS error:', data.type, data.details);
                        if (data.fatal) {
                            switch (data.type) {
                                case Hls.ErrorTypes.NETWORK_ERROR:
                                    console.error('Network error, trying to recover...');
                                    hls.startLoad();
                                    break;
                                case Hls.ErrorTypes.MEDIA_ERROR:
                                    console.error('Media error, trying to recover...');
                                    hls.recoverMediaError();
                                    break;
                                default:
                                    console.error('Fatal error, cannot recover');
                                    hls.destroy();
                                    break;
                            }
                        }
                    });
                    
                    console.log('HLS player initialized with hls.js');
                } else if (video.canPlayType('application/vnd.apple.mpegurl')) {
                    // Safari 原生支持 HLS
                    video.src = src;
                    console.log('Using native HLS support (Safari)');
                } else {
                    console.error('No HLS support in this browser');
                }
            });
        })();
        
        // 侧边栏切换
        function toggleSidebar() {
            const sidebar = document.getElementById('sidebar');
            sidebar.classList.toggle('open');
        }
        
        // 点击内容区域关闭侧边栏（移动端）
        document.addEventListener('click', function(e) {
            const sidebar = document.getElementById('sidebar');
            const menuToggle = document.querySelector('.menu-toggle');
            
            if (window.innerWidth <= 768 && 
                !sidebar.contains(e.target) && 
                !menuToggle.contains(e.target) &&
                sidebar.classList.contains('open')) {
                sidebar.classList.remove('open');
            }
        });
        
        // ESC 键关闭侧边栏
        document.addEventListener('keydown', function(e) {
            if (e.key === 'Escape') {
                const sidebar = document.getElementById('sidebar');
                if (sidebar.classList.contains('open')) {
                    sidebar.classList.remove('open');
                }
            }
        });
        
        // 平滑滚动到锚点
        document.querySelectorAll('a[href^="#"]').forEach(anchor => {
            anchor.addEventListener('click', function (e) {
                e.preventDefault();
                const target = document.querySelector(this.getAttribute('href'));
                if (target) {
                    target.scrollIntoView({
                        behavior: 'smooth',
                        block: 'start'
                    });
                }
            });
        });
        
        // 自动激活当前章节（基于 URL）
        (function() {
            const currentPath = window.location.pathname.split('/').pop();
            if (currentPath) {
                const links = document.querySelectorAll('.nav-link');
                links.forEach(link => {
                    const href = link.getAttribute('href');
                    if (href === currentPath) {
                        link.classList.add('active');
                    } else {
                        link.classList.remove('active');
                    }
                });
            }
        })();
        
        // 媒体播放器浮动功能
        (function() {
            const mediaPlayerContainer = document.getElementById('mediaPlayerContainer');
            const mediaPlayerInline = document.getElementById('mediaPlayerInline');
            const articleContent = document.querySelector('.markdown-body');
            
            if (!mediaPlayerContainer || !articleContent) return;
            
            // 查找内容中的视频或音频元素
            const videoElement = articleContent.querySelector('video');
            const audioElement = articleContent.querySelector('audio');
            
            if (!videoElement && !audioElement) return;
            
            // 获取原始媒体元素
            const originalMedia = videoElement || audioElement;
            const isVideo = !!videoElement;
            
            // 克隆媒体元素到播放器容器
            const clonedMedia = originalMedia.cloneNode(true);
            clonedMedia.id = isVideo ? 'floatingVideo' : 'floatingAudio';
            clonedMedia.classList.add(isVideo ? 'float-video' : 'float-audio');
            
            // 同步播放状态
            function syncMediaState(source, target) {
                if (source && target) {
                    target.currentTime = source.currentTime;
                    if (!source.paused) {
                        target.play().catch(() => {});
                    }
                }
            }
            
            // 显示内联播放器
            mediaPlayerInline.innerHTML = '';
            mediaPlayerInline.appendChild(clonedMedia);
            mediaPlayerContainer.style.display = 'block';
            
            // 隐藏原始播放器但保留占位符
            const placeholder = document.createElement('div');
            placeholder.className = 'media-player-placeholder';
            placeholder.innerHTML = isVideo ? '📺 视频播放器已移至顶部' : ' 音频播放器已移至顶部';
            originalMedia.parentNode.insertBefore(placeholder, originalMedia);
            originalMedia.style.display = 'none';
            
            // 创建浮动播放器
            const floatPlayer = document.createElement('div');
            floatPlayer.className = 'media-player-float hidden';
            floatPlayer.id = 'mediaPlayerFloat';
            
            const closeBtn = document.createElement('button');
            closeBtn.className = 'float-close-btn';
            closeBtn.innerHTML = '×';
            closeBtn.onclick = function() {
                floatPlayer.classList.add('hidden');
            };
            
            const floatMedia = clonedMedia.cloneNode(true);
            floatMedia.id = isVideo ? 'floatVideoPlayer' : 'floatAudioPlayer';
            
            floatPlayer.appendChild(closeBtn);
            floatPlayer.appendChild(floatMedia);
            document.body.appendChild(floatPlayer);
            
            // 同步所有媒体元素的播放状态
            const allMediaElements = [clonedMedia, floatMedia, originalMedia];
            allMediaElements.forEach(media => {
                if (media) {
                    media.addEventListener('play', function() {
                        allMediaElements.forEach(m => {
                            if (m && m !== media && !m.paused) {
                                m.pause();
                            }
                        });
                        syncMediaState(media, clonedMedia);
                        syncMediaState(media, floatMedia);
                    });
                    
                    media.addEventListener('pause', function() {
                        syncMediaState(media, clonedMedia);
                        syncMediaState(media, floatMedia);
                    });
                    
                    media.addEventListener('timeupdate', function() {
                        syncMediaState(media, clonedMedia);
                        syncMediaState(media, floatMedia);
                    });
                }
            });
            
            // 监听滚动事件
            let isFloating = false;
            const observer = new IntersectionObserver(function(entries) {
                entries.forEach(function(entry) {
                    if (!entry.isIntersecting && !isFloating) {
                        // 原始播放器不可见，显示浮动播放器
                        floatPlayer.classList.remove('hidden');
                        isFloating = true;
                        syncMediaState(clonedMedia, floatMedia);
                    } else if (entry.isIntersecting && isFloating) {
                        // 原始播放器可见，隐藏浮动播放器
                        floatPlayer.classList.add('hidden');
                        isFloating = false;
                    }
                });
            }, {
                threshold: 0.1
            });
            
            observer.observe(mediaPlayerContainer);
        })();
    </script>
</body>
</html>
