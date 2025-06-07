// 添加语言支持
const translations = {
    'zh': {
        // 通用
        'app_title': 'Coding Tester',
        'loading': '加载中...',
        'scan_btn': '扫描编程语言',
        'test_btn': '测试通信',
        'search_placeholder': '搜索编程语言...',
        
        // 标签页
        'tab_installed': '已安装',
        'tab_missing': '未安装',
        'tab_packages': '包搜索',
        'tab_ai_assistant': 'AI助手',
        
        // 包搜索
        'package_search_title': '包搜索',
        'package_manager_label': '包管理器:',
        'package_name_label': '包名称:',
        'package_name_placeholder': '输入包名称...',
        'search_btn_text': '搜索',
        'no_results': '没有找到匹配的包',
        'no_input': '请输入包名称',
        'searching': '正在搜索...',
        'install_command': '安装命令:',
        'copy_btn': '复制',
        'view_details': '查看详情',
        
        // 语言卡片
        'installed_status': '已安装',
        'missing_status': '未安装',
        'package_manager': '包管理器:',
        'no_installed_languages': '没有已安装的编程语言',
        'no_missing_languages': '没有未安装的编程语言',
        
        // 语言详情
        'language_details': '详情',
        'status': '状态',
        'version': '版本',
        'missing_deps': '缺失依赖',
        'installed_packages': '已安装的包',
        'recommended_packages': '推荐安装的包',
        'package_tutorial': '包管理器教程',
        'install_package': '安装包',
        'search_package': '搜索包',
        'update_package': '更新包',
        'view_tutorial': '查看详细教程',
        'installation_guide': '安装指南',
        'download': '下载',
        'install_tutorial': '安装教程',
        
        // AI助手
        'ai_settings': 'AI助手设置',
        'ai_provider': 'AI提供商:',
        'api_key': 'API密钥:',
        'api_key_placeholder': '输入API密钥...',
        'custom_endpoint': '自定义端点:',
        'custom_endpoint_placeholder': '输入自定义API端点URL...',
        'save_config': '保存配置',
        'view_docs': '查看API文档',
        'add_detection_data': '添加语言检测数据',
        'detection_data_hint': '将当前检测到的编程语言数据发送给AI分析',
        'ai_welcome': '欢迎使用编程AI助手，请设置API密钥后开始对话。',
        'user_input_placeholder': '输入你的编程问题...',
        'send_btn': '发送',
        'processing_data': '正在处理数据...',
        
        // 系统信息
        'os_info': '操作系统：',
        'arch_info': '架构：',
        'cpu_info': 'CPU 核心：',
        
        // 主题切换器
        'theme_settings': '主题设置',
        'default_theme': '默认主题',
        'dark_theme': '深色主题',
        'teal_theme': '蓝绿主题',
        'purple_theme': '紫色主题',
        'orange_theme': '橙色主题',
        'pink_theme': '粉色主题',
        'green_theme': '绿色主题',
        'yellow_theme': '黄色主题',
        'red_theme': '红色主题',
        'high_contrast': '高对比度',
        'search_theme': '搜索主题...',
        'auto_theme': '自动切换主题',
        'day_theme': '白天主题:',
        'night_theme': '夜间主题:',
        
        // 壁纸选择器
        'wallpaper_settings': '壁纸设置',
        'no_wallpaper': '无壁纸',
        'pattern_1': '格子纹',
        'pattern_2': '圆点纹',
        'pattern_3': '棋盘纹',
        'pattern_4': '网格纹',
        'pattern_5': '渐变',
        'pattern_6': '菱形纹',
        'pattern_7': '细网格',
        'pattern_8': '光晕',
        'pattern_9': '斜纹',
        'pattern_10': '彩虹渐变',
        
        // 设置面板
        'settings': '设置',
        'themes': '主题',
        'wallpapers': '壁纸',
        'auto_settings': '自动',
        'search_settings': '搜索主题或壁纸...'
    },
    'en': {
        // General
        'app_title': 'Coding Tester',
        'loading': 'Loading...',
        'scan_btn': 'Scan Programming Languages',
        'test_btn': 'Test Connection',
        'search_placeholder': 'Search programming languages...',
        
        // Tabs
        'tab_installed': 'Installed',
        'tab_missing': 'Not Installed',
        'tab_packages': 'Package Search',
        'tab_ai_assistant': 'AI Assistant',
        
        // Package search
        'package_search_title': 'Package Search',
        'package_manager_label': 'Package Manager:',
        'package_name_label': 'Package Name:',
        'package_name_placeholder': 'Enter package name...',
        'search_btn_text': 'Search',
        'no_results': 'No matching packages found',
        'no_input': 'Please enter a package name',
        'searching': 'Searching...',
        'install_command': 'Install Command:',
        'copy_btn': 'Copy',
        'view_details': 'View Details',
        
        // Language cards
        'installed_status': 'Installed',
        'missing_status': 'Not Installed',
        'package_manager': 'Package Manager:',
        'no_installed_languages': 'No installed programming languages',
        'no_missing_languages': 'No missing programming languages',
        
        // Language details
        'language_details': 'Details',
        'status': 'Status',
        'version': 'Version',
        'missing_deps': 'Missing Dependencies',
        'installed_packages': 'Installed Packages',
        'recommended_packages': 'Recommended Packages',
        'package_tutorial': 'Package Manager Tutorial',
        'install_package': 'Install Package',
        'search_package': 'Search Package',
        'update_package': 'Update Package',
        'view_tutorial': 'View Detailed Tutorial',
        'installation_guide': 'Installation Guide',
        'download': 'Download',
        'install_tutorial': 'Installation Tutorial',
        
        // AI assistant
        'ai_settings': 'AI Assistant Settings',
        'ai_provider': 'AI Provider:',
        'api_key': 'API Key:',
        'api_key_placeholder': 'Enter API key...',
        'custom_endpoint': 'Custom Endpoint:',
        'custom_endpoint_placeholder': 'Enter custom API endpoint URL...',
        'save_config': 'Save Configuration',
        'view_docs': 'View API Documentation',
        'add_detection_data': 'Add Language Detection Data',
        'detection_data_hint': 'Send detected programming language data to AI for analysis',
        'ai_welcome': 'Welcome to the Programming AI Assistant. Please set up your API key to start a conversation.',
        'user_input_placeholder': 'Enter your programming question...',
        'send_btn': 'Send',
        'processing_data': 'Processing data...',
        
        // System info
        'os_info': 'OS:',
        'arch_info': 'Architecture:',
        'cpu_info': 'CPU Cores:',
        
        // Theme switcher
        'theme_settings': 'Theme Settings',
        'default_theme': 'Default Theme',
        'dark_theme': 'Dark Theme',
        'teal_theme': 'Teal Theme',
        'purple_theme': 'Purple Theme',
        'orange_theme': 'Orange Theme',
        'pink_theme': 'Pink Theme',
        'green_theme': 'Green Theme',
        'yellow_theme': 'Yellow Theme',
        'red_theme': 'Red Theme',
        'high_contrast': 'High Contrast',
        'search_theme': 'Search themes...',
        'auto_theme': 'Auto Theme',
        'day_theme': 'Day Theme:',
        'night_theme': 'Night Theme:',
        
        // Wallpaper switcher
        'wallpaper_settings': 'Wallpaper Settings',
        'no_wallpaper': 'No Wallpaper',
        'pattern_1': 'Checkered',
        'pattern_2': 'Polka Dots',
        'pattern_3': 'Checkerboard',
        'pattern_4': 'Grid',
        'pattern_5': 'Gradient',
        'pattern_6': 'Diamond',
        'pattern_7': 'Fine Grid',
        'pattern_8': 'Glow',
        'pattern_9': 'Diagonal',
        'pattern_10': 'Rainbow',
        
        // Settings panel
        'settings': 'Settings',
        'themes': 'Themes',
        'wallpapers': 'Wallpapers',
        'auto_settings': 'Auto',
        'search_settings': 'Search themes or wallpapers...'
    },
    'ru': {
        // Общие
        'app_title': 'Coding Tester',
        'loading': 'Загрузка...',
        'scan_btn': 'Сканировать языки программирования',
        'test_btn': 'Проверить соединение',
        'search_placeholder': 'Поиск языков программирования...',
        
        // Вкладки
        'tab_installed': 'Установленные',
        'tab_missing': 'Не установленные',
        'tab_packages': 'Поиск пакетов',
        'tab_ai_assistant': 'ИИ-помощник',
        
        // Поиск пакетов
        'package_search_title': 'Поиск пакетов',
        'package_manager_label': 'Менеджер пакетов:',
        'package_name_label': 'Название пакета:',
        'package_name_placeholder': 'Введите название пакета...',
        'search_btn_text': 'Поиск',
        'no_results': 'Подходящие пакеты не найдены',
        'no_input': 'Пожалуйста, введите название пакета',
        'searching': 'Поиск...',
        'install_command': 'Команда установки:',
        'copy_btn': 'Копировать',
        'view_details': 'Подробнее',
        
        // Карточки языков
        'installed_status': 'Установлен',
        'missing_status': 'Не установлен',
        'package_manager': 'Менеджер пакетов:',
        'no_installed_languages': 'Нет установленных языков программирования',
        'no_missing_languages': 'Нет отсутствующих языков программирования',
        
        // Детали языка
        'language_details': 'Сведения о языке',
        'status': 'Статус',
        'version': 'Версия',
        'missing_deps': 'Отсутствующие зависимости',
        'installed_packages': 'Установленные пакеты',
        'recommended_packages': 'Рекомендуемые пакеты',
        'package_tutorial': 'Руководство по менеджеру пакетов',
        'install_package': 'Установить пакет',
        'search_package': 'Найти пакет',
        'update_package': 'Обновить пакет',
        'view_tutorial': 'Просмотреть подробное руководство',
        'installation_guide': 'Руководство по установке',
        'download': 'Скачать',
        'install_tutorial': 'Руководство по установке',
        
        // ИИ-помощник
        'ai_settings': 'Настройки ИИ-помощника',
        'ai_provider': 'Провайдер ИИ:',
        'api_key': 'API-ключ:',
        'api_key_placeholder': 'Введите API-ключ...',
        'custom_endpoint': 'Пользовательский эндпоинт:',
        'custom_endpoint_placeholder': 'Введите URL пользовательского API...',
        'save_config': 'Сохранить настройки',
        'view_docs': 'Просмотреть API-документацию',
        'add_detection_data': 'Добавить данные обнаружения языка',
        'detection_data_hint': 'Отправить обнаруженные данные языков программирования для анализа ИИ',
        'ai_welcome': 'Добро пожаловать в ИИ-помощник по программированию. Пожалуйста, настройте API-ключ, чтобы начать разговор.',
        'user_input_placeholder': 'Введите ваш вопрос по программированию...',
        'send_btn': 'Отправить',
        'processing_data': 'Обработка данных...',
        
        // Системная информация
        'os_info': 'ОС:',
        'arch_info': 'Архитектура:',
        'cpu_info': 'Ядра ЦП:',
        
        // Переключатель тем
        'theme_settings': 'Настройки темы',
        'default_theme': 'Стандартная тема',
        'dark_theme': 'Темная тема',
        'teal_theme': 'Бирюзовая тема',
        'purple_theme': 'Фиолетовая тема',
        'orange_theme': 'Оранжевая тема',
        'pink_theme': 'Розовая тема',
        'green_theme': 'Зеленая тема',
        'yellow_theme': 'Желтая тема',
        'red_theme': 'Красная тема',
        'high_contrast': 'Высокий контраст',
        'search_theme': 'Поиск тем...',
        'auto_theme': 'Автоматическая тема',
        'day_theme': 'Дневная тема:',
        'night_theme': 'Ночная тема:',
        
        // Переключатель обоев
        'wallpaper_settings': 'Настройки обоев',
        'no_wallpaper': 'Без обоев',
        'pattern_1': 'Клетчатый',
        'pattern_2': 'Точечный',
        'pattern_3': 'Шахматный',
        'pattern_4': 'Сетка',
        'pattern_5': 'Градиент',
        'pattern_6': 'Ромбовидный',
        'pattern_7': 'Тонкая сетка',
        'pattern_8': 'Свечение',
        'pattern_9': 'Диагональный',
        'pattern_10': 'Радуга',
        
        // Панель настроек
        'settings': 'Настройки',
        'themes': 'Темы',
        'wallpapers': 'Обои',
        'auto_settings': 'Авто',
        'search_settings': 'Поиск тем или обоев...'
    }
};

// 当前语言
let currentLanguage = 'zh';

// 初始化Wails运行时
window.onload = function() {
    // 等待Wails运行时准备就绪
    window.runtime.WindowSetTitle("Coding Tester");
    
    // 加载语言设置
    loadLanguagePreference();
    
    // 添加语言切换按钮
    addLanguageSwitcher();
    
    // 获取系统信息
    window.go.main.App.GetSystemInfo().then(info => {
        document.getElementById('os-info').textContent = getText('os_info') + ' ' + formatOSName(info.os);
        document.getElementById('arch-info').textContent = getText('arch_info') + ' ' + info.arch;
        document.getElementById('cpu-info').textContent = getText('cpu_info') + ' ' + info.cpus;
    });

    // 设置扫描按钮事件
    document.getElementById('scan-btn').addEventListener('click', scanLanguages);
    
    // 设置测试按钮事件
    document.getElementById('test-btn').addEventListener('click', testBackendConnection);
    
    // 设置搜索按钮事件
    document.getElementById('search-btn').addEventListener('click', searchPackages);
    
    // 设置语言搜索事件
    document.getElementById('language-search').addEventListener('input', searchLanguages);

    // 设置标签页切换
    const tabButtons = document.querySelectorAll('.tab-btn');
    tabButtons.forEach(button => {
        button.addEventListener('click', () => {
            const tabName = button.getAttribute('data-tab');
            
            // 切换活动标签按钮
            tabButtons.forEach(btn => btn.classList.remove('active'));
            button.classList.add('active');
            
            // 切换活动内容
            document.querySelectorAll('.tab-content').forEach(content => {
                content.classList.remove('active');
            });
            document.getElementById(`${tabName}-tab`).classList.add('active');
            
            // 如果有搜索词，重新应用搜索
            const searchInput = document.getElementById('language-search');
            if (searchInput.value.trim() !== '' && tabName !== 'packages' && tabName !== 'ai-assistant') {
                searchLanguages();
            }
            
            // 如果是AI助手标签，初始化AI助手
            if (tabName === 'ai-assistant') {
                initAIAssistant();
            }
        });
    });

    // 设置关闭模态框按钮
    document.getElementById('close-modal').addEventListener('click', () => {
        document.getElementById('language-modal').style.display = 'none';
    });

    // 点击模态框外部关闭
    window.addEventListener('click', (e) => {
        const modal = document.getElementById('language-modal');
        if (e.target === modal) {
            modal.style.display = 'none';
        }
    });
    
    // 回车键触发搜索
    document.getElementById('package-name').addEventListener('keypress', (e) => {
        if (e.key === 'Enter') {
            searchPackages();
        }
    });
    
    // 设置AI相关事件
    setupAIEvents();
    
    // 初始化主题系统
    initThemeSystem();
    
    // 应用当前语言
    applyLanguage();
};

// 添加语言切换按钮
function addLanguageSwitcher() {
    // 创建语言切换器容器
    const languageSwitcher = document.createElement('div');
    languageSwitcher.className = 'language-switcher';
    languageSwitcher.style.display = 'flex';
    languageSwitcher.style.alignItems = 'center';
    languageSwitcher.style.gap = '8px';
    languageSwitcher.style.marginLeft = '20px';
    
    // 添加中文按钮
    const zhButton = document.createElement('button');
    zhButton.textContent = '中文';
    zhButton.className = currentLanguage === 'zh' ? 'lang-btn active' : 'lang-btn';
    zhButton.style.padding = '5px 10px';
    zhButton.style.border = '1px solid #ccc';
    zhButton.style.borderRadius = '4px';
    zhButton.style.cursor = 'pointer';
    if (currentLanguage === 'zh') {
        zhButton.style.backgroundColor = '#4a6cf7';
        zhButton.style.color = 'white';
    }
    zhButton.addEventListener('click', () => {
        if (currentLanguage !== 'zh') {
            currentLanguage = 'zh';
            saveLanguagePreference();
            applyLanguage();
            zhButton.className = 'lang-btn active';
            zhButton.style.backgroundColor = '#4a6cf7';
            zhButton.style.color = 'white';
            enButton.className = 'lang-btn';
            enButton.style.backgroundColor = '';
            enButton.style.color = '';
            ruButton.className = 'lang-btn';
            ruButton.style.backgroundColor = '';
            ruButton.style.color = '';
        }
    });
    
    // 添加英文按钮
    const enButton = document.createElement('button');
    enButton.textContent = 'English';
    enButton.className = currentLanguage === 'en' ? 'lang-btn active' : 'lang-btn';
    enButton.style.padding = '5px 10px';
    enButton.style.border = '1px solid #ccc';
    enButton.style.borderRadius = '4px';
    enButton.style.cursor = 'pointer';
    if (currentLanguage === 'en') {
        enButton.style.backgroundColor = '#4a6cf7';
        enButton.style.color = 'white';
    }
    enButton.addEventListener('click', () => {
        if (currentLanguage !== 'en') {
            currentLanguage = 'en';
            saveLanguagePreference();
            applyLanguage();
            enButton.className = 'lang-btn active';
            enButton.style.backgroundColor = '#4a6cf7';
            enButton.style.color = 'white';
            zhButton.className = 'lang-btn';
            zhButton.style.backgroundColor = '';
            zhButton.style.color = '';
            ruButton.className = 'lang-btn';
            ruButton.style.backgroundColor = '';
            ruButton.style.color = '';
        }
    });
    
    // 添加俄语按钮
    const ruButton = document.createElement('button');
    ruButton.textContent = 'Русский';
    ruButton.className = currentLanguage === 'ru' ? 'lang-btn active' : 'lang-btn';
    ruButton.style.padding = '5px 10px';
    ruButton.style.border = '1px solid #ccc';
    ruButton.style.borderRadius = '4px';
    ruButton.style.cursor = 'pointer';
    if (currentLanguage === 'ru') {
        ruButton.style.backgroundColor = '#4a6cf7';
        ruButton.style.color = 'white';
    }
    ruButton.addEventListener('click', () => {
        if (currentLanguage !== 'ru') {
            currentLanguage = 'ru';
            saveLanguagePreference();
            applyLanguage();
            ruButton.className = 'lang-btn active';
            ruButton.style.backgroundColor = '#4a6cf7';
            ruButton.style.color = 'white';
            zhButton.className = 'lang-btn';
            zhButton.style.backgroundColor = '';
            zhButton.style.color = '';
            enButton.className = 'lang-btn';
            enButton.style.backgroundColor = '';
            enButton.style.color = '';
        }
    });
    
    // 添加按钮到容器
    languageSwitcher.appendChild(zhButton);
    languageSwitcher.appendChild(enButton);
    languageSwitcher.appendChild(ruButton);
    
    // 添加到页面
    const header = document.querySelector('header');
    if (header) {
        // 检查是否已存在语言切换器，如果存在则移除
        const existingLangSwitcher = header.querySelector('.language-switcher');
        if (existingLangSwitcher) {
            header.removeChild(existingLangSwitcher);
        }
        header.appendChild(languageSwitcher);
    }
}

// 获取翻译文本
function getText(key) {
    if (translations[currentLanguage] && translations[currentLanguage][key]) {
        return translations[currentLanguage][key];
    }
    // 如果找不到翻译，返回中文版本或键名
    return translations['zh'][key] || key;
}

// 应用当前语言
function applyLanguage() {
    // 更新标题
    document.title = getText('app_title');
    window.runtime.WindowSetTitle(getText('app_title'));
    
    // 更新按钮文本
    document.getElementById('scan-btn').textContent = getText('scan_btn');
    document.getElementById('test-btn').textContent = getText('test_btn');
    document.getElementById('search-btn').textContent = getText('search_btn_text');
    document.getElementById('send-btn').textContent = getText('send_btn');
    
    // 更新输入框占位符
    document.getElementById('language-search').placeholder = getText('search_placeholder');
    document.getElementById('package-name').placeholder = getText('package_name_placeholder');
    document.getElementById('user-input').placeholder = getText('user_input_placeholder');
    
    // 更新标签页
    document.querySelector('[data-tab="installed"]').innerHTML = getText('tab_installed') + ' <span id="installed-count">0</span>';
    document.querySelector('[data-tab="missing"]').innerHTML = getText('tab_missing') + ' <span id="missing-count">0</span>';
    document.querySelector('[data-tab="packages"]').textContent = getText('tab_packages');
    document.querySelector('[data-tab="ai-assistant"]').textContent = getText('tab_ai_assistant');
    
    // 更新包搜索区域
    document.querySelector('.package-search-container h3').textContent = getText('package_search_title');
    document.querySelector('label[for="package-manager"]').textContent = getText('package_manager_label');
    document.querySelector('label[for="package-name"]').textContent = getText('package_name_label');
    
    // 更新AI助手区域
    document.querySelector('.ai-settings h3').textContent = getText('ai_settings');
    document.querySelector('label[for="ai-provider"]').textContent = getText('ai_provider');
    document.querySelector('label[for="api-key"]').textContent = getText('api_key');
    document.getElementById('api-key').placeholder = getText('api_key_placeholder');
    document.querySelector('label[for="custom-endpoint"]').textContent = getText('custom_endpoint');
    document.getElementById('custom-endpoint').placeholder = getText('custom_endpoint_placeholder');
    document.getElementById('save-ai-config-btn').textContent = getText('save_config');
    document.getElementById('provider-docs-link').textContent = getText('view_docs');
    document.getElementById('add-detection-data-btn').textContent = getText('add_detection_data');
    document.querySelector('.hint-text').textContent = getText('detection_data_hint');
    
    // 更新系统信息标签
    const osLabel = document.querySelector('#os-info').previousElementSibling;
    const archLabel = document.querySelector('#arch-info').previousElementSibling;
    const cpuLabel = document.querySelector('#cpu-info').previousElementSibling;
    
    if (osLabel) osLabel.textContent = getText('os_info');
    if (archLabel) archLabel.textContent = getText('arch_info');
    if (cpuLabel) cpuLabel.textContent = getText('cpu_info');
    
    // 更新页脚 - 移除页脚文本
    const footerText = document.querySelector('footer p');
    if (footerText) {
        footerText.textContent = '';
    }
    
    // 更新主题切换器标题
    document.querySelectorAll('.theme-option').forEach(option => {
        const theme = option.getAttribute('data-theme');
        switch (theme) {
            case 'default':
                option.title = getText('default_theme');
                break;
            case 'dark':
                option.title = getText('dark_theme');
                break;
            case 'teal':
                option.title = getText('teal_theme');
                break;
            case 'purple':
                option.title = getText('purple_theme');
                break;
            case 'orange':
                option.title = getText('orange_theme');
                break;
            case 'pink':
                option.title = getText('pink_theme');
                break;
            case 'green':
                option.title = getText('green_theme');
                break;
            case 'yellow':
                option.title = getText('yellow_theme');
                break;
            case 'red':
                option.title = getText('red_theme');
                break;
            case 'high-contrast':
                option.title = getText('high_contrast');
                break;
        }
    });
    
    // 更新主题切换按钮标题
    const themeToggleBtn = document.querySelector('.theme-toggle-btn');
    if (themeToggleBtn) {
        themeToggleBtn.title = getText('theme_settings');
    }
    
    // 如果有系统消息，更新它
    const systemMessage = document.querySelector('.message.system .message-content p');
    if (systemMessage && systemMessage.textContent === translations['zh']['ai_welcome']) {
        systemMessage.textContent = getText('ai_welcome');
    }
    
    // 更新动态生成的元素
    // 更新下载按钮
    document.querySelectorAll('.download-btn').forEach(btn => {
        // 检测当前按钮文本是否为其他语言的"下载"
        if (btn.textContent.includes('下载') || 
            btn.textContent.includes('Download') || 
            btn.textContent.includes('Скачать')) {
            // 提取语言名称（假设格式是"下载 语言名"）
            const langName = btn.textContent.split(' ').slice(1).join(' ');
            btn.textContent = `${getText('download')} ${langName}`;
        }
    });
    
    // 更新安装教程按钮
    document.querySelectorAll('.tutorial-btn').forEach(btn => {
        if (btn.textContent === '安装教程' || 
            btn.textContent === 'Installation Tutorial' || 
            btn.textContent === 'Руководство по установке') {
            btn.textContent = getText('install_tutorial');
        } else if (btn.textContent === '查看详细教程' || 
                   btn.textContent === 'View Detailed Tutorial' || 
                   btn.textContent === 'Просмотреть подробное руководство') {
            btn.textContent = getText('view_tutorial');
        }
    });
    
    // 更新复制按钮
    document.querySelectorAll('.copy-btn').forEach(btn => {
        if (btn.textContent.trim() === '复制' || 
            btn.textContent.trim() === 'Copy' || 
            btn.textContent.trim() === 'Копировать') {
            btn.textContent = getText('copy_btn');
        }
    });
    
    // 更新查看详情链接
    document.querySelectorAll('.download-link').forEach(link => {
        if (link.textContent.trim() === '查看详情' || 
            link.textContent.trim() === 'View Details' || 
            link.textContent.trim() === 'Подробнее') {
            link.textContent = getText('view_details');
        }
    });
    
    // 更新命令标签
    document.querySelectorAll('.command-label').forEach(label => {
        if (label.textContent.includes('安装命令') || 
            label.textContent.includes('Install Command') || 
            label.textContent.includes('Команда установки')) {
            label.textContent = getText('install_command');
        }
    });
    
    // 更新模态框中的文本
    const modalTitle = document.getElementById('modal-title');
    if (modalTitle) {
        const titleText = modalTitle.textContent;
        // 如果标题包含语言名称和"详情"字样
        if (titleText.includes('详情') || 
            titleText.includes('Details') || 
            titleText.includes('Сведения')) {
            // 提取语言名称（假设格式是"语言名 详情"）
            const langName = titleText.split(' ')[0];
            modalTitle.textContent = `${langName} ${getText('language_details')}`;
        }
    }
    
    // 刷新当前视图
    refreshCurrentView();
}

// 刷新当前视图
function refreshCurrentView() {
    // 获取当前活动的标签页
    const activeTab = document.querySelector('.tab-btn.active');
    if (activeTab) {
        const tabName = activeTab.getAttribute('data-tab');
        
        // 根据当前标签页刷新视图
        if (tabName === 'installed' || tabName === 'missing') {
            // 重新渲染语言卡片
            const container = document.getElementById(`${tabName}-languages`);
            const cards = container.querySelectorAll('.language-card');
            
            if (cards.length === 0) {
                container.innerHTML = `<p class="no-results">${getText(tabName === 'installed' ? 'no_installed_languages' : 'no_missing_languages')}</p>`;
            } else {
                cards.forEach(card => {
                    const statusElement = card.querySelector('.status');
                    if (statusElement) {
                        statusElement.textContent = getText(tabName === 'installed' ? 'installed_status' : 'missing_status');
                    }
                    
                    const packageManagerElement = card.querySelector('.package-manager');
                    if (packageManagerElement) {
                        const packageManagerText = packageManagerElement.textContent;
                        const packageManagerName = packageManagerText.split(':')[1].trim();
                        packageManagerElement.textContent = `${getText('package_manager')} ${packageManagerName}`;
                    }
                });
            }
        } else if (tabName === 'packages') {
            // 更新搜索结果中的文本
            const searchResults = document.getElementById('search-results');
            const noResults = searchResults.querySelector('.no-results');
            if (noResults) {
                if (noResults.textContent === translations['zh']['no_results']) {
                    noResults.textContent = getText('no_results');
                } else if (noResults.textContent === translations['zh']['no_input']) {
                    noResults.textContent = getText('no_input');
                }
            }
            
            // 更新搜索结果中的安装命令和查看详情按钮
            const resultItems = searchResults.querySelectorAll('.result-item');
            resultItems.forEach(item => {
                const commandLabel = item.querySelector('.command-label');
                if (commandLabel) {
                    commandLabel.textContent = getText('install_command');
                }
                
                const copyBtn = item.querySelector('.copy-btn');
                if (copyBtn) {
                    copyBtn.textContent = getText('copy_btn');
                }
                
                const downloadLink = item.querySelector('.download-link');
                if (downloadLink) {
                    downloadLink.textContent = getText('view_details');
                }
            });
        }
    }
    
    // 更新模态框标题
    const modalTitle = document.getElementById('modal-title');
    if (modalTitle && modalTitle.textContent.includes('详情')) {
        const langName = modalTitle.textContent.split(' ')[0];
        modalTitle.textContent = `${langName} ${getText('language_details')}`;
    }
}

// 保存语言偏好
function saveLanguagePreference() {
    localStorage.setItem('preferred-language', currentLanguage);
    
    // 同时保存到后端
    window.go.main.App.SaveLanguageConfig({
        language: currentLanguage
    }).catch(err => {
        console.error("保存语言配置失败:", err);
    });
}

// 加载语言偏好
function loadLanguagePreference() {
    // 首先尝试从后端加载
    window.go.main.App.GetLanguageConfig()
        .then(config => {
            if (config && config.language && (config.language === 'zh' || config.language === 'en' || config.language === 'ru')) {
                currentLanguage = config.language;
            } else {
                // 如果后端没有配置，尝试从本地存储加载
                const savedLanguage = localStorage.getItem('preferred-language');
                if (savedLanguage && (savedLanguage === 'zh' || savedLanguage === 'en' || savedLanguage === 'ru')) {
                    currentLanguage = savedLanguage;
                    // 同步到后端
                    window.go.main.App.SaveLanguageConfig({
                        language: currentLanguage
                    }).catch(err => {
                        console.error("同步语言配置失败:", err);
                    });
                } else {
                    // 根据浏览器语言设置默认语言
                    const browserLang = navigator.language.toLowerCase();
                    if (browserLang.startsWith('zh')) {
                        currentLanguage = 'zh';
                    } else if (browserLang.startsWith('ru')) {
                        currentLanguage = 'ru';
                    } else {
                        currentLanguage = 'en';
                    }
                    // 保存默认语言到后端
                    window.go.main.App.SaveLanguageConfig({
                        language: currentLanguage
                    }).catch(err => {
                        console.error("保存默认语言配置失败:", err);
                    });
                }
            }
        })
        .catch(err => {
            console.error("加载语言配置失败:", err);
            // 出错时使用本地存储或浏览器语言
            const savedLanguage = localStorage.getItem('preferred-language');
            if (savedLanguage && (savedLanguage === 'zh' || savedLanguage === 'en' || savedLanguage === 'ru')) {
                currentLanguage = savedLanguage;
            } else {
                // 根据浏览器语言设置默认语言
                const browserLang = navigator.language.toLowerCase();
                if (browserLang.startsWith('zh')) {
                    currentLanguage = 'zh';
                } else if (browserLang.startsWith('ru')) {
                    currentLanguage = 'ru';
                } else {
                    currentLanguage = 'en';
                }
            }
        });
}

// 初始化主题系统
function initThemeSystem() {
    console.log('初始化主题系统');
    
    // 加载保存的主题
    loadSavedTheme();
    
    // 设置主题切换事件
    const themeOptions = document.querySelectorAll('.theme-option');
    themeOptions.forEach(option => {
        // 点击事件 - 永久切换主题
        option.addEventListener('click', () => {
            const theme = option.getAttribute('data-theme');
            console.log('主题切换: 点击了主题', theme);
            
            // 更新活动状态
            themeOptions.forEach(opt => {
                opt.classList.remove('active');
                console.log('移除活动状态:', opt.getAttribute('data-theme'));
            });
            option.classList.add('active');
            console.log('添加活动状态:', theme);
            
            // 设置主题
            setTheme(theme);
            console.log('应用主题:', theme);
            
            // 保存主题选择
            saveTheme(theme);
            
            // 添加动画效果
            option.classList.add('previewing');
            setTimeout(() => {
                option.classList.remove('previewing');
            }, 500);
        });
        
        // 鼠标悬停事件 - 临时预览主题
        option.addEventListener('mouseenter', () => {
            // 如果已经是活动主题，不需要预览
            if (option.classList.contains('active')) {
                console.log('鼠标悬停: 已是活动主题，跳过预览');
                return;
            }
            
            // 保存当前主题
            const currentTheme = document.body.getAttribute('data-theme') || 'default';
            option.dataset.previousTheme = currentTheme;
            console.log('鼠标悬停: 保存当前主题', currentTheme);
            
            // 临时应用新主题
            const previewTheme = option.getAttribute('data-theme');
            setTheme(previewTheme, true);
            console.log('鼠标悬停: 预览主题', previewTheme);
            
            // 添加预览指示
            option.classList.add('previewing');
        });
        
        // 鼠标离开事件 - 恢复原主题
        option.addEventListener('mouseleave', () => {
            // 如果是活动主题，不需要恢复
            if (option.classList.contains('active')) {
                console.log('鼠标离开: 已是活动主题，不恢复');
                option.classList.remove('previewing');
                return;
            }
            
            // 如果有保存的主题，恢复它
            if (option.dataset.previousTheme) {
                const previousTheme = option.dataset.previousTheme;
                console.log('鼠标离开: 恢复之前的主题', previousTheme);
                setTheme(previousTheme, true);
                delete option.dataset.previousTheme;
            }
            
            // 移除预览指示
            option.classList.remove('previewing');
        });
    });
    
    // 添加主题切换器的折叠/展开功能
    const themeSwitcher = document.getElementById('theme-switcher');
    if (themeSwitcher) {
        // 获取折叠/展开按钮
        const toggleButton = themeSwitcher.querySelector('.theme-toggle-btn');
        
        if (toggleButton) {
            // 设置按钮标题
            toggleButton.title = getText('theme_settings');
            
            // 默认折叠状态
            let isCollapsed = localStorage.getItem('theme-switcher-collapsed') === 'true';
            
            // 应用初始状态
            if (isCollapsed) {
                themeSwitcher.classList.add('collapsed');
            }
            
            // 点击事件
            toggleButton.addEventListener('click', () => {
                isCollapsed = !isCollapsed;
                if (isCollapsed) {
                    themeSwitcher.classList.add('collapsed');
                } else {
                    themeSwitcher.classList.remove('collapsed');
                }
                localStorage.setItem('theme-switcher-collapsed', isCollapsed);
            });
        }
    }
    
    // 初始化自动主题功能
    initAutoTheme();
    
    // 根据系统偏好设置初始主题
    if (!localStorage.getItem('preferred-theme')) {
        // 检测系统偏好
        if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
            setTheme('dark');
            document.querySelector('.theme-option[data-theme="dark"]').classList.add('active');
        } else {
            setTheme('default');
            document.querySelector('.theme-option[data-theme="default"]').classList.add('active');
        }
    }
    
    // 监听系统主题变化
    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
            if (!localStorage.getItem('preferred-theme')) {
                setTheme(e.matches ? 'dark' : 'default');
                
                // 更新活动状态
                document.querySelectorAll('.theme-option').forEach(opt => opt.classList.remove('active'));
                document.querySelector(`.theme-option[data-theme="${e.matches ? 'dark' : 'default'}"]`).classList.add('active');
            }
        });
    }
    
    // 添加主题搜索功能
    addThemeSearch();
}

// 初始化自动主题功能
function initAutoTheme() {
    // 检查是否启用了自动主题
    const autoThemeEnabled = localStorage.getItem('auto-theme-enabled') === 'true';
    
    if (autoThemeEnabled) {
        // 设置基于时间的主题
        setTimeBasedTheme();
        
        // 每小时检查一次
        setInterval(setTimeBasedTheme, 60 * 60 * 1000);
    }
}

// 设置基于时间的主题
function setTimeBasedTheme() {
    const hour = new Date().getHours();
    
    // 早上6点到晚上6点使用浅色主题，其他时间使用深色主题
    if (hour >= 6 && hour < 18) {
        // 白天主题
        const dayTheme = localStorage.getItem('day-theme') || 'default';
        setTheme(dayTheme);
        updateActiveThemeOption(dayTheme);
    } else {
        // 夜间主题
        const nightTheme = localStorage.getItem('night-theme') || 'dark';
        setTheme(nightTheme);
        updateActiveThemeOption(nightTheme);
    }
}

// 更新活动主题选项
function updateActiveThemeOption(theme) {
    const themeOptions = document.querySelectorAll('.theme-option');
    themeOptions.forEach(opt => opt.classList.remove('active'));
    
    const activeOption = document.querySelector(`.theme-option[data-theme="${theme}"]`);
    if (activeOption) {
        activeOption.classList.add('active');
    }
}

// 添加主题搜索功能
function addThemeSearch() {
    const themeSwitcher = document.getElementById('theme-switcher');
    if (!themeSwitcher) return;
    
    // 创建搜索输入框
    const searchInput = document.createElement('input');
    searchInput.type = 'text';
    searchInput.className = 'theme-search';
    searchInput.placeholder = getText('search_theme');
    searchInput.style.display = 'none'; // 默认隐藏
    searchInput.style.gridColumn = 'span 3';
    searchInput.style.padding = '5px';
    searchInput.style.borderRadius = '4px';
    searchInput.style.border = '1px solid var(--light-gray)';
    searchInput.style.marginBottom = '5px';
    
    // 创建搜索按钮
    const searchButton = document.createElement('button');
    searchButton.className = 'theme-search-btn';
    searchButton.innerHTML = '🔍';
    searchButton.title = getText('search_theme');
    searchButton.style.width = '30px';
    searchButton.style.height = '30px';
    searchButton.style.borderRadius = '50%';
    searchButton.style.backgroundColor = 'var(--secondary-color)';
    searchButton.style.color = 'white';
    searchButton.style.border = 'none';
    searchButton.style.cursor = 'pointer';
    searchButton.style.position = 'absolute';
    searchButton.style.top = '15px';
    searchButton.style.left = '15px';
    
    // 添加搜索按钮点击事件
    searchButton.addEventListener('click', () => {
        const isVisible = searchInput.style.display !== 'none';
        searchInput.style.display = isVisible ? 'none' : 'block';
        if (!isVisible) {
            searchInput.focus();
        }
    });
    
    // 添加搜索输入事件
    searchInput.addEventListener('input', () => {
        const searchTerm = searchInput.value.toLowerCase();
        const themeOptions = document.querySelectorAll('.theme-option');
        
        themeOptions.forEach(option => {
            const themeName = option.title.toLowerCase();
            if (searchTerm === '' || themeName.includes(searchTerm)) {
                option.style.display = 'block';
            } else {
                option.style.display = 'none';
            }
        });
    });
    
    // 添加到主题切换器
    themeSwitcher.prepend(searchButton);
    themeSwitcher.querySelector('.theme-options-grid').prepend(searchInput);
}

// 添加自动主题设置菜单
function addAutoThemeSettings() {
    const themeSwitcher = document.getElementById('theme-switcher');
    if (!themeSwitcher) return;
    
    // 创建设置按钮
    const settingsButton = document.createElement('button');
    settingsButton.className = 'theme-settings-btn';
    settingsButton.innerHTML = '⚙️';
    settingsButton.title = getText('theme_settings');
    settingsButton.style.width = '30px';
    settingsButton.style.height = '30px';
    settingsButton.style.borderRadius = '50%';
    settingsButton.style.backgroundColor = 'var(--gray-color)';
    settingsButton.style.color = 'white';
    settingsButton.style.border = 'none';
    settingsButton.style.cursor = 'pointer';
    settingsButton.style.position = 'absolute';
    settingsButton.style.top = '15px';
    settingsButton.style.right = '15px';
    
    // 创建设置菜单
    const settingsMenu = document.createElement('div');
    settingsMenu.className = 'theme-settings-menu';
    settingsMenu.style.display = 'none';
    settingsMenu.style.position = 'absolute';
    settingsMenu.style.top = '50px';
    settingsMenu.style.right = '10px';
    settingsMenu.style.backgroundColor = 'var(--card-bg)';
    settingsMenu.style.borderRadius = 'var(--border-radius)';
    settingsMenu.style.boxShadow = 'var(--box-shadow)';
    settingsMenu.style.padding = '15px';
    settingsMenu.style.zIndex = '1001';
    settingsMenu.style.width = '200px';
    
    // 自动主题开关
    const autoThemeEnabled = localStorage.getItem('auto-theme-enabled') === 'true';
    const autoThemeToggle = document.createElement('div');
    autoThemeToggle.className = 'settings-item';
    autoThemeToggle.style.marginBottom = '10px';
    
    const autoThemeLabel = document.createElement('label');
    autoThemeLabel.style.display = 'flex';
    autoThemeLabel.style.alignItems = 'center';
    autoThemeLabel.style.cursor = 'pointer';
    
    const autoThemeCheckbox = document.createElement('input');
    autoThemeCheckbox.type = 'checkbox';
    autoThemeCheckbox.checked = autoThemeEnabled;
    autoThemeCheckbox.style.marginRight = '10px';
    
    autoThemeLabel.appendChild(autoThemeCheckbox);
    autoThemeLabel.appendChild(document.createTextNode(getText('auto_theme')));
    autoThemeToggle.appendChild(autoThemeLabel);
    
    // 白天主题选择
    const dayThemeSelect = document.createElement('div');
    dayThemeSelect.className = 'settings-item';
    dayThemeSelect.style.marginBottom = '10px';
    
    const dayThemeLabel = document.createElement('label');
    dayThemeLabel.textContent = getText('day_theme');
    dayThemeLabel.style.display = 'block';
    dayThemeLabel.style.marginBottom = '5px';
    
    const dayThemeDropdown = document.createElement('select');
    dayThemeDropdown.style.width = '100%';
    dayThemeDropdown.style.padding = '5px';
    
    // 夜间主题选择
    const nightThemeSelect = document.createElement('div');
    nightThemeSelect.className = 'settings-item';
    
    const nightThemeLabel = document.createElement('label');
    nightThemeLabel.textContent = getText('night_theme');
    nightThemeLabel.style.display = 'block';
    nightThemeLabel.style.marginBottom = '5px';
    
    const nightThemeDropdown = document.createElement('select');
    nightThemeDropdown.style.width = '100%';
    nightThemeDropdown.style.padding = '5px';
    
    // 添加主题选项
    const themes = [
        { value: 'default', text: getText('default_theme') },
        { value: 'dark', text: getText('dark_theme') },
        { value: 'teal', text: getText('teal_theme') },
        { value: 'purple', text: getText('purple_theme') },
        { value: 'orange', text: getText('orange_theme') },
        { value: 'pink', text: getText('pink_theme') },
        { value: 'green', text: getText('green_theme') },
        { value: 'yellow', text: getText('yellow_theme') },
        { value: 'red', text: getText('red_theme') },
        { value: 'high-contrast', text: getText('high_contrast') }
    ];
    
    themes.forEach(theme => {
        const dayOption = document.createElement('option');
        dayOption.value = theme.value;
        dayOption.textContent = theme.text;
        dayThemeDropdown.appendChild(dayOption);
        
        const nightOption = document.createElement('option');
        nightOption.value = theme.value;
        nightOption.textContent = theme.text;
        nightThemeDropdown.appendChild(nightOption);
    });
    
    // 设置默认选中值
    dayThemeDropdown.value = localStorage.getItem('day-theme') || 'default';
    nightThemeDropdown.value = localStorage.getItem('night-theme') || 'dark';
    
    // 添加事件监听器
    autoThemeCheckbox.addEventListener('change', () => {
        localStorage.setItem('auto-theme-enabled', autoThemeCheckbox.checked);
        dayThemeSelect.style.display = autoThemeCheckbox.checked ? 'block' : 'none';
        nightThemeSelect.style.display = autoThemeCheckbox.checked ? 'block' : 'none';
        
        if (autoThemeCheckbox.checked) {
            initAutoTheme();
        }
    });
    
    dayThemeDropdown.addEventListener('change', () => {
        localStorage.setItem('day-theme', dayThemeDropdown.value);
        if (autoThemeCheckbox.checked) {
            setTimeBasedTheme();
        }
    });
    
    nightThemeDropdown.addEventListener('change', () => {
        localStorage.setItem('night-theme', nightThemeDropdown.value);
        if (autoThemeCheckbox.checked) {
            setTimeBasedTheme();
        }
    });
    
    // 组装设置菜单
    dayThemeSelect.appendChild(dayThemeLabel);
    dayThemeSelect.appendChild(dayThemeDropdown);
    nightThemeSelect.appendChild(nightThemeLabel);
    nightThemeSelect.appendChild(nightThemeDropdown);
    
    settingsMenu.appendChild(autoThemeToggle);
    settingsMenu.appendChild(dayThemeSelect);
    settingsMenu.appendChild(nightThemeSelect);
    
    // 初始隐藏白天/夜间主题选择
    dayThemeSelect.style.display = autoThemeEnabled ? 'block' : 'none';
    nightThemeSelect.style.display = autoThemeEnabled ? 'block' : 'none';
    
    // 添加设置按钮点击事件
    settingsButton.addEventListener('click', (e) => {
        e.stopPropagation();
        settingsMenu.style.display = settingsMenu.style.display === 'none' ? 'block' : 'none';
    });
    
    // 点击其他地方关闭设置菜单
    document.addEventListener('click', (e) => {
        if (!settingsMenu.contains(e.target) && e.target !== settingsButton) {
            settingsMenu.style.display = 'none';
        }
    });
    
    // 添加到主题切换器
    themeSwitcher.appendChild(settingsButton);
    themeSwitcher.appendChild(settingsMenu);
    
    // 监听语言变化，更新UI文本
    document.addEventListener('languagechange', updateAutoThemeText);
}

// 更新自动主题设置的文本
function updateAutoThemeText() {
    const settingsButton = document.querySelector('.theme-settings-btn');
    if (settingsButton) {
        settingsButton.title = getText('theme_settings');
    }
    
    const searchButton = document.querySelector('.theme-search-btn');
    if (searchButton) {
        searchButton.title = getText('search_theme');
    }
    
    const searchInput = document.querySelector('.theme-search');
    if (searchInput) {
        searchInput.placeholder = getText('search_theme');
    }
    
    const autoThemeLabel = document.querySelector('.settings-item label');
    if (autoThemeLabel && autoThemeLabel.lastChild.nodeType === Node.TEXT_NODE) {
        autoThemeLabel.lastChild.textContent = getText('auto_theme');
    }
    
    const dayThemeLabel = document.querySelector('.settings-item:nth-child(2) label');
    if (dayThemeLabel) {
        dayThemeLabel.textContent = getText('day_theme');
    }
    
    const nightThemeLabel = document.querySelector('.settings-item:nth-child(3) label');
    if (nightThemeLabel) {
        nightThemeLabel.textContent = getText('night_theme');
    }
    
    // 更新主题选项
    const dayThemeDropdown = document.querySelector('.settings-item:nth-child(2) select');
    const nightThemeDropdown = document.querySelector('.settings-item:nth-child(3) select');
    
    if (dayThemeDropdown && nightThemeDropdown) {
        const themes = [
            { value: 'default', text: getText('default_theme') },
            { value: 'dark', text: getText('dark_theme') },
            { value: 'teal', text: getText('teal_theme') },
            { value: 'purple', text: getText('purple_theme') },
            { value: 'orange', text: getText('orange_theme') },
            { value: 'pink', text: getText('pink_theme') },
            { value: 'green', text: getText('green_theme') },
            { value: 'yellow', text: getText('yellow_theme') },
            { value: 'red', text: getText('red_theme') },
            { value: 'high-contrast', text: getText('high_contrast') }
        ];
        
        // 保存当前选择的值
        const daySelected = dayThemeDropdown.value;
        const nightSelected = nightThemeDropdown.value;
        
        // 清空并重新填充选项
        dayThemeDropdown.innerHTML = '';
        nightThemeDropdown.innerHTML = '';
        
        themes.forEach(theme => {
            const dayOption = document.createElement('option');
            dayOption.value = theme.value;
            dayOption.textContent = theme.text;
            dayThemeDropdown.appendChild(dayOption);
            
            const nightOption = document.createElement('option');
            nightOption.value = theme.value;
            nightOption.textContent = theme.text;
            nightThemeDropdown.appendChild(nightOption);
        });
        
        // 恢复选择
        dayThemeDropdown.value = daySelected;
        nightThemeDropdown.value = nightSelected;
    }
}

// 在语言变化时更新主题相关文本
document.addEventListener('languagechange', () => {
    // 更新主题选项标题
    document.querySelectorAll('.theme-option').forEach(option => {
        const theme = option.getAttribute('data-theme');
        switch (theme) {
            case 'default':
                option.title = getText('default_theme');
                break;
            case 'dark':
                option.title = getText('dark_theme');
                break;
            case 'teal':
                option.title = getText('teal_theme');
                break;
            case 'purple':
                option.title = getText('purple_theme');
                break;
            case 'orange':
                option.title = getText('orange_theme');
                break;
            case 'pink':
                option.title = getText('pink_theme');
                break;
            case 'green':
                option.title = getText('green_theme');
                break;
            case 'yellow':
                option.title = getText('yellow_theme');
                break;
            case 'red':
                option.title = getText('red_theme');
                break;
            case 'high-contrast':
                option.title = getText('high_contrast');
                break;
        }
    });
    
    // 更新主题切换按钮标题
    const themeToggleBtn = document.querySelector('.theme-toggle-btn');
    if (themeToggleBtn) {
        themeToggleBtn.title = getText('theme_settings');
    }
    
    // 更新自动主题设置文本
    updateAutoThemeText();
});

// 设置主题
function setTheme(theme, isPreview = false) {
    console.log('setTheme 函数被调用，主题:', theme);
    
    // 移除所有主题
    document.body.removeAttribute('data-theme');
    console.log('移除了现有主题属性');
    
    // 设置新主题
    if (theme !== 'default') {
        document.body.setAttribute('data-theme', theme);
        console.log('设置了新主题属性:', theme);
    } else {
        console.log('使用默认主题，不设置data-theme属性');
    }
    
    // 应用主题特定的调整
    applyThemeAdjustments(theme);
    
    // 如果不是预览，触发主题变更事件
    if (!isPreview) {
        // 创建自定义事件
        const event = new CustomEvent('themechange', { detail: { theme } });
        document.dispatchEvent(event);
        console.log('触发了themechange事件');
    }
}

// 应用主题特定调整
function applyThemeAdjustments(theme) {
    // 例如，对于深色主题，可能需要调整某些元素
    if (theme === 'dark') {
        // 调整图表颜色等
    } else if (theme === 'high-contrast') {
        // 增加边框等以提高可访问性
    }
}

// 保存主题选择
function saveTheme(theme) {
    // 本地存储
    localStorage.setItem('preferred-theme', theme);
    
    // 保存到后端
    window.go.main.App.SaveThemeConfig({theme: theme})
        .then(() => {
            console.log('主题配置已保存到后端');
        })
        .catch(error => {
            console.error('保存主题配置失败:', error);
        });
}

// 加载保存的主题
function loadSavedTheme() {
    console.log('加载保存的主题');
    // 首先尝试从后端加载
    window.go.main.App.GetThemeConfig()
        .then(config => {
            console.log('从后端加载主题配置:', config);
            if (config && config.theme) {
                setTheme(config.theme);
                
                // 更新活动状态
                document.querySelectorAll('.theme-option').forEach(opt => opt.classList.remove('active'));
                const activeOption = document.querySelector(`.theme-option[data-theme="${config.theme}"]`);
                if (activeOption) {
                    activeOption.classList.add('active');
                }
                
                // 同步到本地存储
                localStorage.setItem('preferred-theme', config.theme);
            } else {
                // 如果后端没有配置，尝试从本地存储加载
                const savedTheme = localStorage.getItem('preferred-theme');
                if (savedTheme) {
                    setTheme(savedTheme);
                    
                    // 更新活动状态
                    document.querySelectorAll('.theme-option').forEach(opt => opt.classList.remove('active'));
                    const activeOption = document.querySelector(`.theme-option[data-theme="${savedTheme}"]`);
                    if (activeOption) {
                        activeOption.classList.add('active');
                    }
                    
                    // 保存到后端
                    window.go.main.App.SaveThemeConfig({theme: savedTheme}).catch(console.error);
                }
            }
        })
        .catch(error => {
            console.error('加载主题配置失败:', error);
            
            // 如果后端加载失败，尝试从本地存储加载
            const savedTheme = localStorage.getItem('preferred-theme');
            if (savedTheme) {
                setTheme(savedTheme);
                
                // 更新活动状态
                document.querySelectorAll('.theme-option').forEach(opt => opt.classList.remove('active'));
                const activeOption = document.querySelector(`.theme-option[data-theme="${savedTheme}"]`);
                if (activeOption) {
                    activeOption.classList.add('active');
                }
            }
        });
}

// 测试后端连接
function testBackendConnection() {
    const resultElement = document.getElementById('test-result');
    resultElement.textContent = "正在测试...";
    
    window.go.main.App.TestFunction()
        .then(result => {
            resultElement.textContent = result;
            console.log("后端连接测试成功:", result);
        })
        .catch(error => {
            resultElement.textContent = "测试失败: " + error;
            console.error("后端连接测试失败:", error);
        });
}

// 格式化操作系统名称
function formatOSName(os) {
    switch (os) {
        case 'windows':
            return 'Windows';
        case 'darwin':
            return 'macOS';
        case 'linux':
            return 'Linux';
        default:
            return os;
    }
}

// 扫描语言
function scanLanguages() {
    const scanBtn = document.getElementById('scan-btn');
    const loadingSpinner = document.getElementById('loading-spinner');
    const resultsContainer = document.getElementById('results-container');
    const progressBar = document.getElementById('scan-progress-bar');
    const progressText = document.getElementById('scan-progress-text');
    
    // 重置进度条
    if (progressBar && progressText) {
        progressBar.style.width = '0%';
        progressText.textContent = '0%';
    }
    
    // 显示加载动画
    scanBtn.style.display = 'none';
    loadingSpinner.style.display = 'flex';
    resultsContainer.style.display = 'none';
    
    // 更新加载提示文本
    const loadingText = document.querySelector('#loading-spinner p');
    if (loadingText) {
        loadingText.textContent = '正在扫描系统中的编程语言...';
    }
    
    // 改进的进度条逻辑 - 使用更平滑的曲线和动态时间间隔
    let progress = 0;
    let lastProgress = 0;
    let updateInterval = 100; // 初始更新间隔更短，提供更流畅的体验
    let stallCounter = 0; // 用于检测进度停滞
    let timeoutCounter = 0; // 用于检测总体超时
    let maxProgress = 99.5; // 提高最大进度上限，避免在98%停留太久
    
    // 预定义的关键进度点，确保进度条不会在某些点停留太久
    const keyProgressPoints = [
        { point: 30, message: '正在检测基础编程语言...' },
        { point: 50, message: '正在检测扩展语言...' },
        { point: 70, message: '正在分析包管理器...' },
        { point: 85, message: '正在收集依赖信息...' },
        { point: 95, message: '正在整理检测结果...' },
        { point: 99, message: '即将完成...' }
    ];
    
    // 更新进度条和文本的函数
    function updateProgressUI(currentProgress, message) {
        if (progressBar && progressText) {
            // 使用小数点后一位的精度显示进度，提供更精细的反馈
            const displayProgress = Math.round(currentProgress * 10) / 10;
            progressBar.style.width = `${displayProgress}%`;
            progressText.textContent = `${displayProgress}%`;
            
            // 当进度超过50%时，改变文字颜色以提高可读性
            if (displayProgress > 50) {
                progressText.style.color = 'white';
            } else {
                progressText.style.color = 'var(--dark-color)';
            }
            
            // 如果提供了消息，更新加载提示文本
            if (message && loadingText) {
                loadingText.textContent = message;
            }
        }
    }
    
    const progressInterval = setInterval(() => {
        // 计算新的进度增量，使用非线性曲线
        let increment;
        let message = null;
        
        // 检查是否到达了关键进度点
        for (const keyPoint of keyProgressPoints) {
            if (progress < keyPoint.point && progress + 1 >= keyPoint.point) {
                message = keyPoint.message;
                // 在关键点稍微加速
                increment = 1.0;
                break;
            }
        }
        
        // 如果不是关键点，使用正常的增量计算
        if (!increment) {
            if (progress < 30) {
                // 开始阶段快速增长
                increment = 0.8 + Math.random() * 1.2;
            } else if (progress < 60) {
                // 中间阶段中等速度
                increment = 0.5 + Math.random() * 0.8;
            } else if (progress < 85) {
                // 后期减慢
                increment = 0.3 + Math.random() * 0.5;
            } else if (progress < 95) {
                // 接近完成时更慢，但不要太慢
                increment = 0.2 + Math.random() * 0.3;
            } else {
                // 最后阶段，保持一定速度以避免停滞感
                increment = 0.1 + Math.random() * 0.2;
            }
        }
        
        // 检测进度是否停滞
        if (Math.round(progress * 10) === Math.round(lastProgress * 10)) {
            stallCounter++;
            
            // 如果在同一百分比停留超过4次，强制增加进度
            if (stallCounter > 4) {
                increment = 0.5 + Math.random() * 0.5; // 更大的强制增量
                stallCounter = 0;
            }
        } else {
            stallCounter = 0;
        }
        
        // 总体超时检测 - 如果进度更新次数过多，加速进度
        timeoutCounter++;
        if (timeoutCounter > 100 && progress < 90) {
            // 如果更新了100次还没到90%，加速进度
            increment *= 2;
        }
        
        lastProgress = progress;
        progress += increment;
        
        // 限制最大进度，留一点空间给完成时设置100%
        progress = Math.min(progress, maxProgress);
        
        // 更新UI
        updateProgressUI(progress, message);
        
        // 动态调整更新间隔，随着进度增加而延长，但保持一定的响应性
        if (progress > 90 && updateInterval < 250) {
            clearInterval(progressInterval);
            updateInterval = 250;
            startProgressInterval();
        } else if (progress > 70 && updateInterval < 200) {
            clearInterval(progressInterval);
            updateInterval = 200;
            startProgressInterval();
        } else if (progress > 40 && updateInterval < 150) {
            clearInterval(progressInterval);
            updateInterval = 150;
            startProgressInterval();
        }
    }, updateInterval);
    
    // 辅助函数，用于重新启动进度间隔
    function startProgressInterval() {
        progressInterval = setInterval(arguments.callee.caller, updateInterval);
    }
    
    // 调用后端检测语言
    window.go.main.App.DetectLanguages().then(languages => {
        // 停止进度模拟
        clearInterval(progressInterval);
        
        // 设置进度为100%
        updateProgressUI(100, '扫描完成！');
        
        // 更新加载提示
        if (loadingText) {
            loadingText.textContent = '扫描完成，正在处理结果...';
        }
        
        setTimeout(() => {
            // 处理结果
            const installedLanguages = languages.filter(lang => lang.installed);
            const missingLanguages = languages.filter(lang => !lang.installed);
            
            // 更新计数
            document.getElementById('installed-count').textContent = installedLanguages.length;
            document.getElementById('missing-count').textContent = missingLanguages.length;
            
            // 渲染已安装的语言
            renderLanguages(installedLanguages, 'installed-languages', true);
            
            // 渲染未安装的语言
            renderLanguages(missingLanguages, 'missing-languages', false);
            
            // 隐藏加载动画，显示结果
            scanBtn.style.display = 'block';
            loadingSpinner.style.display = 'none';
            resultsContainer.style.display = 'block';
            
            // 显示扫描结果统计
            if (installedLanguages.length > 0 || missingLanguages.length > 0) {
                const totalMsg = `扫描完成！已检测到 ${installedLanguages.length} 种已安装语言和 ${missingLanguages.length} 种未安装语言。`;
                const resultElement = document.getElementById('test-result');
                if (resultElement) {
                    resultElement.textContent = totalMsg;
                    resultElement.style.color = 'var(--secondary-color)';
                }
            }
        }, 800); // 增加延迟，让用户能够看到100%的进度
    }).catch(error => {
        // 停止进度模拟
        clearInterval(progressInterval);
        
        console.error('扫描语言时出错:', error);
        scanBtn.style.display = 'block';
        loadingSpinner.style.display = 'none';
        
        const resultElement = document.getElementById('test-result');
        if (resultElement) {
            resultElement.textContent = '扫描语言时出错，请重试。错误: ' + error.message;
            resultElement.style.color = 'var(--danger-color)';
        } else {
            alert('扫描语言时出错，请重试。');
        }
    });
}

// 渲染语言卡片
function renderLanguages(languages, containerId, isInstalled) {
    const container = document.getElementById(containerId);
    container.innerHTML = '';
    
    if (languages.length === 0) {
        container.innerHTML = `<p class="no-results">没有${isInstalled ? '已安装' : '未安装'}的编程语言</p>`;
        return;
    }
    
    // 对语言按名称排序
    languages.sort((a, b) => a.name.localeCompare(b.name));
    
    languages.forEach(lang => {
        const card = document.createElement('div');
        card.className = `language-card ${isInstalled ? 'installed' : 'missing'}`;
        
        let versionText = isInstalled ? 
            `<div class="version">${truncateVersion(lang.version)}</div>` : 
            '';
        
        let packageInfo = '';
        if (isInstalled && lang.packageManager) {
            packageInfo = `<div class="package-manager">包管理器: ${lang.packageManager}</div>`;
        }
        
        card.innerHTML = `
            <h3>${lang.name}</h3>
            ${versionText}
            ${packageInfo}
            <span class="status ${isInstalled ? 'installed' : 'missing'}">
                ${isInstalled ? '已安装' : '未安装'}
            </span>
        `;
        
        // 添加点击事件，显示详情
        card.addEventListener('click', () => {
            showLanguageDetails(lang, isInstalled);
        });
        
        container.appendChild(card);
    });
}

// 截断版本字符串
function truncateVersion(version) {
    if (!version) return '';
    
    // 如果版本字符串太长，截断它
    if (version.length > 50) {
        return version.substring(0, 50) + '...';
    }
    
    return version;
}

// 获取包管理器教程
async function getPackageTutorials() {
    try {
        return await window.go.main.App.GetPackageTutorials();
    } catch (error) {
        console.error('获取包管理器教程出错:', error);
        return [];
    }
}

// 获取特定语言的包列表
async function getLanguagePackages(languageName) {
    try {
        const packages = await window.go.main.App.GetLanguagePackages(languageName);
        return packages;
    } catch (error) {
        console.error(`获取${languageName}包列表出错:`, error);
        return [];
    }
}

// 显示语言详情
async function showLanguageDetails(language, isInstalled) {
    const modal = document.getElementById('language-modal');
    const modalTitle = document.getElementById('modal-title');
    const modalBody = document.getElementById('modal-body');
    
    modalTitle.textContent = `${language.name} ${getText('language_details')}`;
    
    // 获取包管理器教程
    const tutorials = await getPackageTutorials();
    const packageTutorial = tutorials.find(t => t.name.includes(language.packageManager));
    
    let content = `
        <div class="language-detail">
            <div class="detail-item">
                <span class="label">${getText('status')}:</span>
                <span class="value">
                    <span class="status ${isInstalled ? 'installed' : 'missing'}">
                        ${isInstalled ? getText('installed_status') : getText('missing_status')}
                    </span>
                </span>
            </div>
    `;
    
    if (isInstalled) {
        content += `
            <div class="detail-item">
                <span class="label">${getText('version')}:</span>
                <span class="value">${language.version || '未知'}</span>
            </div>
        `;
        
        if (language.packageManager) {
            content += `
                <div class="detail-item">
                    <span class="label">${getText('package_manager')}:</span>
                    <span class="value">${language.packageManager}</span>
                </div>
            `;
        }
        
        if (language.missingDeps && language.missingDeps.length > 0) {
            content += `
                <div class="detail-item">
                    <span class="label">${getText('missing_deps')}:</span>
                    <div class="missing-deps">
                        ${language.missingDeps.map(dep => `<span class="dep-item">${dep}</span>`).join('')}
                    </div>
                </div>
            `;
        }
        
        // 显示已安装的包
        if (language.packages && language.packages.length > 0) {
            content += `
                <div class="detail-item packages-section">
                    <span class="label">${getText('installed_packages')} (${language.packages.length}):</span>
                    <div class="packages-list">
                        ${language.packages.map(pkg => `
                            <div class="package-item">
                                <span class="package-name">${pkg.name}</span>
                                <span class="package-version">${pkg.version || ''}</span>
                            </div>
                        `).join('')}
                    </div>
                </div>
            `;
        }
        
        // 显示缺少的推荐包
        const missingPackages = await getMissingPackages(language.name);
        if (missingPackages && missingPackages.length > 0) {
            content += `
                <div class="detail-item packages-section recommended-packages">
                    <span class="label">${getText('recommended_packages')} (${missingPackages.length}):</span>
                    <div class="packages-list">
                        ${missingPackages.map(pkg => `
                            <div class="package-item">
                                <div class="package-info">
                                    <span class="package-name">${pkg.name}</span>
                                    <span class="package-version">${pkg.version || ''}</span>
                                </div>
                                <div class="package-description">${pkg.description || ''}</div>
                            </div>
                        `).join('')}
                    </div>
                </div>
            `;
        }
        
        // 显示包管理器教程
        if (packageTutorial) {
            content += `
                <div class="detail-item package-tutorial">
                    <span class="label">${getText('package_tutorial')}:</span>
                    <div class="tutorial-content">
                        <div class="tutorial-item">
                            <span class="tutorial-label">${getText('install_package')}:</span>
                            <code>${packageTutorial.installCmd}</code>
                        </div>
                        <div class="tutorial-item">
                            <span class="tutorial-label">${getText('search_package')}:</span>
                            <code>${packageTutorial.searchCmd}</code>
                        </div>
                        <div class="tutorial-item">
                            <span class="tutorial-label">${getText('update_package')}:</span>
                            <code>${packageTutorial.updateCmd}</code>
                        </div>
                        <div class="tutorial-item">
                            <a href="${packageTutorial.tutorialURL}" class="tutorial-btn" target="_blank">${getText('view_tutorial')}</a>
                        </div>
                    </div>
                </div>
            `;
        }
    } else {
        content += `
            <div class="download-section">
                <h3>${getText('installation_guide')}</h3>
                <a href="${language.downloadURL}" class="download-btn" target="_blank">${getText('download')} ${language.name}</a>
                <a href="${language.installTutorial}" class="tutorial-btn" target="_blank">${getText('install_tutorial')}</a>
            </div>
        `;
        
        // 显示推荐包
        if (language.recommendedPkgs && language.recommendedPkgs.length > 0) {
            content += `
                <div class="detail-item packages-section recommended-packages">
                    <span class="label">${getText('recommended_packages')} (${language.recommendedPkgs.length}):</span>
                    <div class="packages-list">
                        ${language.recommendedPkgs.map(pkg => `
                            <div class="package-item">
                                <div class="package-info">
                                    <span class="package-name">${pkg.name}</span>
                                    <span class="package-version">${pkg.version || ''}</span>
                                </div>
                                <div class="package-description">${pkg.description || ''}</div>
                            </div>
                        `).join('')}
                    </div>
                </div>
            `;
        }
    }
    
    content += '</div>';
    
    modalBody.innerHTML = content;
    modal.style.display = 'block';
}

// 获取缺少的包
async function getMissingPackages(languageName) {
    try {
        const missingPackages = await window.go.main.App.GetMissingPackages(languageName);
        return missingPackages;
    } catch (error) {
        console.error(`获取${languageName}缺少的包出错:`, error);
        return [];
    }
}

// 搜索包
async function searchPackages() {
    const packageManager = document.getElementById('package-manager').value;
    const packageName = document.getElementById('package-name').value.trim();
    const searchResults = document.getElementById('search-results');
    
    if (!packageName) {
        searchResults.innerHTML = `<p class="no-results">${getText('no_input')}</p>`;
        return;
    }
    
    // 显示加载中
    searchResults.innerHTML = `
        <div class="loading">
            <div class="spinner"></div>
            <p>${getText('searching')} ${packageName}...</p>
        </div>
    `;
    
    try {
        const results = await window.go.main.App.SearchPackage(packageManager, packageName);
        
        if (results.length === 0) {
            searchResults.innerHTML = `<p class="no-results">${getText('no_results')} "${packageName}"</p>`;
            return;
        }
        
        // 渲染搜索结果
        let html = '';
        results.forEach(pkg => {
            html += `
                <div class="result-item">
                    <div class="result-header">
                        <span class="result-name">${pkg.name}</span>
                        ${pkg.version ? `<span class="result-version">${pkg.version}</span>` : ''}
                    </div>
                    ${pkg.description ? `<p class="result-description">${pkg.description}</p>` : ''}
                    <div class="result-actions">
                        ${pkg.installLink ? `
                            <div class="install-command">
                                <span class="command-label">${getText('install_command')}</span>
                                <code>${pkg.installLink}</code>
                                <button class="copy-btn" onclick="navigator.clipboard.writeText('${pkg.installLink.replace(/'/g, "\\'")}')">
                                    ${getText('copy_btn')}
                                </button>
                            </div>
                        ` : ''}
                        ${pkg.downloadURL ? `
                            <a href="${pkg.downloadURL}" class="download-link" target="_blank">
                                ${getText('view_details')}
                            </a>
                        ` : ''}
                    </div>
                </div>
            `;
        });
        
        searchResults.innerHTML = html;
    } catch (error) {
        console.error('搜索包时出错:', error);
        searchResults.innerHTML = `<p class="no-results">${getText('searching')} ${error.message || '未知错误'}</p>`;
    }
}

// 搜索语言
function searchLanguages(event) {
    const searchTerm = document.getElementById('language-search').value.toLowerCase();
    const activeTab = document.querySelector('.tab-btn.active').getAttribute('data-tab');
    const languageCards = document.querySelectorAll(`#${activeTab}-languages .language-card`);
    
    languageCards.forEach(card => {
        const languageName = card.querySelector('h3').textContent.toLowerCase();
        if (languageName.includes(searchTerm) || searchTerm === '') {
            card.style.display = 'block';
        } else {
            card.style.display = 'none';
        }
    });
}

// 创建简单的Logo
function createLogo() {
    const canvas = document.createElement('canvas');
    canvas.width = 200;
    canvas.height = 200;
    const ctx = canvas.getContext('2d');
    
    // 绘制背景圆
    ctx.fillStyle = '#4a6bff';
    ctx.beginPath();
    ctx.arc(100, 100, 80, 0, Math.PI * 2);
    ctx.fill();
    
    // 绘制内部圆
    ctx.fillStyle = 'white';
    ctx.beginPath();
    ctx.arc(100, 100, 60, 0, Math.PI * 2);
    ctx.fill();
    
    // 绘制网络连接线
    ctx.strokeStyle = '#4a6bff';
    ctx.lineWidth = 8;
    
    // 绘制连接线
    ctx.beginPath();
    ctx.moveTo(70, 70);
    ctx.lineTo(130, 130);
    ctx.stroke();
    
    ctx.beginPath();
    ctx.moveTo(130, 70);
    ctx.lineTo(70, 130);
    ctx.stroke();
    
    // 绘制节点
    ctx.fillStyle = '#34c759';
    ctx.beginPath();
    ctx.arc(70, 70, 12, 0, Math.PI * 2);
    ctx.fill();
    
    ctx.fillStyle = '#ff3b30';
    ctx.beginPath();
    ctx.arc(130, 130, 12, 0, Math.PI * 2);
    ctx.fill();
    
    ctx.fillStyle = '#ffcc00';
    ctx.beginPath();
    ctx.arc(130, 70, 12, 0, Math.PI * 2);
    ctx.fill();
    
    ctx.fillStyle = '#5856d6';
    ctx.beginPath();
    ctx.arc(70, 130, 12, 0, Math.PI * 2);
    ctx.fill();
    
    return canvas.toDataURL();
}

// AI助手相关功能

// 设置AI相关事件
function setupAIEvents() {
    // 保存AI配置按钮
    const saveConfigBtn = document.getElementById('save-ai-config-btn');
    if (saveConfigBtn) {
        saveConfigBtn.addEventListener('click', saveAIConfig);
    }
    
    // 发送消息按钮
    const sendBtn = document.getElementById('send-btn');
    if (sendBtn) {
        sendBtn.addEventListener('click', sendAIMessage);
    }
    
    // 用户输入按下Enter键发送消息
    const userInput = document.getElementById('user-input');
    if (userInput) {
        userInput.addEventListener('keydown', (e) => {
            // Ctrl+Enter发送消息
            if (e.key === 'Enter' && e.ctrlKey) {
                e.preventDefault();
                sendAIMessage();
            }
        });
    }
    
    // AI提供商选择变化时更新UI
    const providerSelect = document.getElementById('ai-provider');
    if (providerSelect) {
        providerSelect.addEventListener('change', updateProviderUI);
    }
    
    // 添加语言检测数据按钮
    const addDetectionDataBtn = document.getElementById('add-detection-data-btn');
    if (addDetectionDataBtn) {
        addDetectionDataBtn.addEventListener('click', addDetectionDataToChat);
    }
}

// 初始化AI助手
async function initAIAssistant() {
    try {
        // 获取AI提供商列表
        const providers = await window.go.main.App.GetAIProviders();
        const providerSelect = document.getElementById('ai-provider');
        
        if (providerSelect) {
            // 清空现有选项
            providerSelect.innerHTML = '';
            
            // 添加提供商选项
            providers.forEach(provider => {
                const option = document.createElement('option');
                option.value = provider.id;
                option.textContent = provider.name;
                option.dataset.docUrl = provider.documentUrl;
                providerSelect.appendChild(option);
            });
            
            // 获取保存的配置
            const config = await window.go.main.App.GetAIConfig();
            
            // 设置选中的提供商
            if (config.selectedProviderId) {
                providerSelect.value = config.selectedProviderId;
            }
            
            // 设置API密钥
            const apiKeyInput = document.getElementById('api-key');
            if (apiKeyInput && config.apiKey) {
                apiKeyInput.value = config.apiKey;
            }
            
            // 设置自定义端点
            const customEndpointInput = document.getElementById('custom-endpoint');
            if (customEndpointInput && config.customEndpoint) {
                customEndpointInput.value = config.customEndpoint;
            }
            
            // 更新UI
            updateProviderUI();
        }
    } catch (error) {
        console.error('初始化AI助手失败:', error);
        showSystemMessage('初始化AI助手失败: ' + error.message);
    }
}

// 更新提供商UI
function updateProviderUI() {
    const providerSelect = document.getElementById('ai-provider');
    const customEndpointContainer = document.getElementById('custom-endpoint-container');
    const providerDocsLink = document.getElementById('provider-docs-link');
    
    if (providerSelect && customEndpointContainer) {
        const selectedProvider = providerSelect.value;
        
        // 显示/隐藏自定义端点输入
        if (selectedProvider === 'custom' || selectedProvider === 'azure_openai') {
            customEndpointContainer.style.display = 'block';
        } else {
            customEndpointContainer.style.display = 'none';
        }
        
        // 更新文档链接
        if (providerDocsLink) {
            const selectedOption = providerSelect.options[providerSelect.selectedIndex];
            const docUrl = selectedOption.dataset.docUrl;
            
            if (docUrl) {
                providerDocsLink.href = docUrl;
                providerDocsLink.style.display = 'inline-block';
            } else {
                providerDocsLink.style.display = 'none';
            }
        }
    }
}

// 保存AI配置
async function saveAIConfig() {
    const providerSelect = document.getElementById('ai-provider');
    const apiKeyInput = document.getElementById('api-key');
    const customEndpointInput = document.getElementById('custom-endpoint');
    
    if (!providerSelect || !apiKeyInput) {
        return;
    }
    
    const selectedProvider = providerSelect.value;
    const apiKey = apiKeyInput.value.trim();
    const customEndpoint = customEndpointInput ? customEndpointInput.value.trim() : '';
    
    if (!apiKey) {
        showSystemMessage('请输入API密钥');
        return;
    }
    
    if ((selectedProvider === 'custom' || selectedProvider === 'azure_openai') && !customEndpoint) {
        showSystemMessage('请输入自定义端点URL');
        return;
    }
    
    try {
        // 保存配置
        const config = {
            selectedProviderId: selectedProvider,
            apiKey: apiKey,
            customEndpoint: customEndpoint
        };
        
        await window.go.main.App.SaveAIConfig(config);
        showSystemMessage('配置已保存');
    } catch (error) {
        console.error('保存AI配置失败:', error);
        showSystemMessage('保存AI配置失败: ' + error.message);
    }
}

// 发送AI消息
async function sendAIMessage() {
    const userInput = document.getElementById('user-input');
    const chatMessages = document.getElementById('chat-messages');
    
    if (!userInput || !chatMessages) {
        return;
    }
    
    const query = userInput.value.trim();
    
    if (!query) {
        return;
    }
    
    // 获取配置
    const config = await window.go.main.App.GetAIConfig();
    
    if (!config.apiKey) {
        showSystemMessage('请先设置API密钥');
        return;
    }
    
    // 添加用户消息
    addMessage('user', query);
    
    // 清空输入框
    userInput.value = '';
    
    // 显示加载中
    const loadingId = showLoadingMessage();
    
    try {
        // 发送请求
        const response = await window.go.main.App.QueryAI(
            config.selectedProviderId,
            config.apiKey,
            config.customEndpoint,
            query
        );
        
        // 移除加载中
        removeLoadingMessage(loadingId);
        
        if (response.success) {
            // 添加AI响应
            addMessage('ai', response.content, response.provider);
        } else {
            // 显示错误
            showSystemMessage('请求失败: ' + response.error);
        }
    } catch (error) {
        // 移除加载中
        removeLoadingMessage(loadingId);
        
        console.error('AI请求失败:', error);
        showSystemMessage('AI请求失败: ' + error.message);
    }
}

// 添加语言检测数据到聊天
async function addDetectionDataToChat() {
    try {
        // 获取进度条元素
        const progressContainer = document.getElementById('ai-progress-container');
        const progressBar = document.getElementById('ai-progress-bar');
        const progressText = document.getElementById('ai-progress-text');
        const progressStatus = document.getElementById('ai-progress-status');
        
        // 重置并显示进度条
        if (progressContainer && progressBar && progressText && progressStatus) {
            progressBar.style.width = '0%';
            progressText.textContent = '0%';
            progressStatus.textContent = '正在初始化数据收集...';
            progressContainer.style.display = 'block';
            
            // 更新进度文字颜色
            progressText.style.color = 'var(--dark-color)';
        }
        
        // 检查是否有检测数据
        const installedLangElements = document.querySelectorAll('#installed-languages .language-card');
        if (installedLangElements.length === 0) {
            // 隐藏进度条
            if (progressContainer) {
                progressContainer.style.display = 'none';
            }
            
            showSystemMessage('请先在"已安装"选项卡中扫描编程语言');
            return;
        }
        
        // 显示正在收集数据的提示
        showSystemMessage('正在收集编程语言数据...');
        
        // 定义数据收集阶段的关键点，每个阶段都有明确的进度值和描述
        const dataCollectionStages = [
            { progress: 3, message: '正在初始化...', delay: 200 },
            { progress: 8, message: '准备收集系统信息...', delay: 300 },
            { progress: 15, message: '正在收集系统信息...', delay: 400 },
            { progress: 22, message: '准备获取编程语言数据...', delay: 250 },
            { progress: 30, message: '正在获取已安装的编程语言...', delay: 500 },
            { progress: 40, message: '正在分析语言数据...', delay: 350 },
            { progress: 50, message: '正在筛选语言数据...', delay: 300 },
            { progress: 60, message: '正在整理语言信息数据...', delay: 400 },
            { progress: 70, message: '正在分析包信息...', delay: 350 },
            { progress: 80, message: '正在处理语言依赖关系...', delay: 300 },
            { progress: 88, message: '正在整合数据...', delay: 300 },
            { progress: 94, message: '数据收集完成，准备发送...', delay: 400 }
        ];
        
        // 顺序执行每个阶段
        for (const stage of dataCollectionStages) {
            updateProgress(progressBar, progressText, progressStatus, stage.progress, stage.message);
            await new Promise(resolve => setTimeout(resolve, stage.delay));
        }
        
        // 获取系统信息
        const systemInfo = {
            os: document.getElementById('os-info').textContent,
            arch: document.getElementById('arch-info').textContent,
            cpus: document.getElementById('cpu-info').textContent
        };
        
        // 获取已安装的语言列表
        const languages = await window.go.main.App.DetectLanguages();
        const installedLanguages = languages.filter(lang => lang.installed);
        
        if (installedLanguages.length === 0) {
            // 隐藏进度条
            if (progressContainer) {
                progressContainer.style.display = 'none';
            }
            
            showSystemMessage('未检测到已安装的编程语言');
            return;
        }
        
        // 显示正在整理数据的提示
        showSystemMessage('正在整理语言信息数据...');
        
        // 准备发送到AI的系统概述
        let summaryText = `系统信息: 操作系统=${systemInfo.os}, 架构=${systemInfo.arch}, CPU核心=${systemInfo.cpus}\n\n`;
        summaryText += `检测到已安装的编程语言(${installedLanguages.length}个):\n`;
        
        // 添加语言信息，最多显示前10个
        const displayLanguages = installedLanguages.slice(0, 10);
        displayLanguages.forEach(lang => {
            summaryText += `- ${lang.name} ${lang.version || ''}\n`;
            if (lang.packageManager) {
                summaryText += `  包管理器: ${lang.packageManager}\n`;
            }
            
            // 添加安装的包信息（如果有且不超过5个）
            if (lang.packages && lang.packages.length > 0) {
                summaryText += `  已安装包(${lang.packages.length}个): `;
                const packagesToShow = lang.packages.slice(0, 5);
                summaryText += packagesToShow.map(pkg => pkg.name).join(', ');
                if (lang.packages.length > 5) {
                    summaryText += ` 等${lang.packages.length}个包`;
                }
                summaryText += '\n';
            }
        });
        
        if (installedLanguages.length > 10) {
            summaryText += `...以及其他${installedLanguages.length - 10}种语言\n`;
        }
        
        // 添加提示信息
        summaryText += "\n请根据我的编程环境，推荐适合我的编程项目和学习路径。";
        
        // 显示成功收集数据的提示
        showSystemMessage('已成功收集编程环境数据并添加到对话中');
        
        // 将系统信息添加到聊天
        addMessage('user', summaryText);
        
        // 自动发送查询
        const config = await window.go.main.App.GetAIConfig();
        
        if (!config.apiKey) {
            // 隐藏进度条
            if (progressContainer) {
                progressContainer.style.display = 'none';
            }
            
            showSystemMessage('请先设置API密钥');
            return;
        }
        
        // 改进的等待AI响应的进度条动画
        updateProgress(progressBar, progressText, progressStatus, 95, '正在发送数据到AI进行分析...');
        
        // 显示加载中
        showSystemMessage('正在发送数据到AI进行分析...');
        const loadingId = showLoadingMessage();
        
        // 改进的进度条动画，使用非线性增长和更小的增量
        let fakeProgress = 95;
        let stallCounter = 0;
        let lastFakeProgress = 95;
        let timeoutCounter = 0;
        
        const fakeProgressInterval = setInterval(() => {
            // 非线性增长，越接近100%增长越慢，但保持最小速度
            const remaining = 99.8 - fakeProgress;
            let increment = Math.max(0.05, remaining * 0.03 * Math.random());
            
            // 超时检测 - 如果更新次数过多，加速进度
            timeoutCounter++;
            if (timeoutCounter > 50) {
                increment = Math.max(0.1, increment * 1.5);
            }
            
            // 防止停滞
            if (Math.round(fakeProgress * 10) === Math.round(lastFakeProgress * 10)) {
                stallCounter++;
                if (stallCounter > 8) {
                    increment = 0.2; // 更大的强制增量
                    stallCounter = 0;
                    
                    // 更新状态文本，提供视觉反馈
                    const waitingMessages = [
                        '正在等待AI分析结果...',
                        'AI正在处理您的数据...',
                        'AI正在生成建议...',
                        '正在优化分析结果...',
                        '即将完成分析...'
                    ];
                    const randomMessage = waitingMessages[Math.floor(Math.random() * waitingMessages.length)];
                    updateProgress(progressBar, progressText, progressStatus, fakeProgress, randomMessage);
                }
            } else {
                stallCounter = 0;
            }
            
            lastFakeProgress = fakeProgress;
            fakeProgress += increment;
            
            if (fakeProgress > 99.8) {
                fakeProgress = 99.8;
            }
            
            updateProgress(progressBar, progressText, progressStatus, fakeProgress, '正在等待AI分析结果...');
        }, 300);
        
        try {
            // 发送请求
            const response = await window.go.main.App.QueryAI(
                config.selectedProviderId,
                config.apiKey,
                config.customEndpoint,
                summaryText
            );
            
            // 停止假进度更新
            clearInterval(fakeProgressInterval);
            
            // 更新进度到100%
            updateProgress(progressBar, progressText, progressStatus, 100, '分析完成！');
            
            // 移除加载中
            removeLoadingMessage(loadingId);
            
            // 隐藏进度条（延迟一会，让用户看到100%）
            setTimeout(() => {
                if (progressContainer) {
                    progressContainer.style.display = 'none';
                }
            }, 1000);
            
            if (response.success) {
                // 添加AI响应
                addMessage('ai', response.content, response.provider);
                showSystemMessage('分析完成，AI已提供针对您编程环境的建议');
            } else {
                // 显示错误
                showSystemMessage('请求失败: ' + response.error);
            }
        } catch (error) {
            // 停止假进度更新
            clearInterval(fakeProgressInterval);
            
            // 隐藏进度条
            if (progressContainer) {
                progressContainer.style.display = 'none';
            }
            
            // 移除加载中
            removeLoadingMessage(loadingId);
            
            console.error('AI请求失败:', error);
            showSystemMessage('AI请求失败: ' + error.message);
        }
    } catch (error) {
        // 隐藏进度条
        const progressContainer = document.getElementById('ai-progress-container');
        if (progressContainer) {
            progressContainer.style.display = 'none';
        }
        
        console.error('添加语言检测数据失败:', error);
        showSystemMessage('添加语言检测数据失败: ' + error.message);
    }
}

// 更新进度条
function updateProgress(progressBar, progressText, progressStatus, progress, statusText) {
    if (!progressBar || !progressText) return;
    
    // 更新进度条 - 添加平滑过渡动画
    // 根据进度值调整过渡时间，高进度值时使用更短的过渡时间，避免感觉停滞
    const transitionDuration = progress > 90 ? 0.3 : (progress > 70 ? 0.4 : 0.5);
    progressBar.style.transition = `width ${transitionDuration}s cubic-bezier(0.4, 0.0, 0.2, 1)`;
    progressBar.style.width = `${progress}%`;
    
    // 使用小数点后一位的精度显示进度，提供更精细的反馈
    const displayProgress = Math.round(progress * 10) / 10;
    
    // 添加动画效果，当进度变化超过1%时，短暂高亮显示进度文本
    const prevProgress = progressText.dataset.lastProgress ? parseFloat(progressText.dataset.lastProgress) : 0;
    if (Math.abs(displayProgress - prevProgress) >= 1.0) {
        progressText.classList.add('progress-highlight');
        setTimeout(() => {
            progressText.classList.remove('progress-highlight');
        }, 300);
    }
    
    // 保存当前进度值以便下次比较
    progressText.dataset.lastProgress = displayProgress;
    progressText.textContent = `${displayProgress}%`;
    
    // 当进度超过50%时，改变文字颜色以提高可读性
    if (progress > 50) {
        progressText.style.color = 'white';
    } else {
        progressText.style.color = 'var(--dark-color)';
    }
    
    // 更新状态文本，添加淡入效果
    if (progressStatus && statusText) {
        // 如果状态文本发生变化，添加淡入效果
        if (progressStatus.textContent !== statusText) {
            progressStatus.style.opacity = '0';
            setTimeout(() => {
                progressStatus.textContent = statusText;
                progressStatus.style.opacity = '1';
            }, 150);
        } else {
            progressStatus.textContent = statusText;
        }
    }
    
    // 当进度接近完成时，添加脉冲效果
    if (progress > 95) {
        progressBar.classList.add('near-complete');
    } else {
        progressBar.classList.remove('near-complete');
    }
}

// 添加消息
function addMessage(type, content, provider = '') {
    const chatMessages = document.getElementById('chat-messages');
    
    if (!chatMessages) {
        return;
    }
    
    const messageDiv = document.createElement('div');
    messageDiv.className = `message ${type}`;
    
    const contentDiv = document.createElement('div');
    contentDiv.className = 'message-content';
    
    // 处理Markdown格式的内容（简单处理代码块）
    let formattedContent = content;
    
    // 处理代码块 ```code```
    formattedContent = formattedContent.replace(/```([\s\S]*?)```/g, function(match, code) {
        return `<pre><code>${code}</code></pre>`;
    });
    
    // 处理内联代码 `code`
    formattedContent = formattedContent.replace(/`([^`]+)`/g, '<code>$1</code>');
    
    // 处理换行
    formattedContent = formattedContent.replace(/\n/g, '<br>');
    
    // 设置内容
    contentDiv.innerHTML = formattedContent;
    
    // 如果是AI消息并有提供商信息，添加提供商标签
    if (type === 'ai' && provider) {
        const providerLabel = document.createElement('div');
        providerLabel.className = 'provider-label';
        providerLabel.textContent = `由 ${provider} 提供`;
        providerLabel.style.fontSize = '0.8em';
        providerLabel.style.color = '#888';
        providerLabel.style.marginTop = '5px';
        contentDiv.appendChild(providerLabel);
    }
    
    messageDiv.appendChild(contentDiv);
    chatMessages.appendChild(messageDiv);
    
    // 滚动到底部
    chatMessages.scrollTop = chatMessages.scrollHeight;
}

// 显示系统消息
function showSystemMessage(message) {
    addMessage('system', message);
}

// 显示加载中消息
function showLoadingMessage() {
    const chatMessages = document.getElementById('chat-messages');
    
    if (!chatMessages) {
        return null;
    }
    
    const id = 'loading-' + Date.now();
    
    const messageDiv = document.createElement('div');
    messageDiv.className = 'message ai';
    messageDiv.id = id;
    
    const contentDiv = document.createElement('div');
    contentDiv.className = 'message-content loading-message';
    
    // 添加加载动画
    const loadingDots = document.createElement('div');
    loadingDots.className = 'loading-dots';
    loadingDots.innerHTML = '<span></span><span></span><span></span>';
    
    contentDiv.appendChild(loadingDots);
    messageDiv.appendChild(contentDiv);
    chatMessages.appendChild(messageDiv);
    
    // 滚动到底部
    chatMessages.scrollTop = chatMessages.scrollHeight;
    
    return id;
}

// 移除加载中消息
function removeLoadingMessage(id) {
    if (!id) {
        return;
    }
    
    const loadingMessage = document.getElementById(id);
    
    if (loadingMessage) {
        loadingMessage.remove();
    }
}

// 在文档加载完成后添加自动主题设置
document.addEventListener('DOMContentLoaded', () => {
    // 延迟一点添加设置菜单，确保主题切换器已初始化
    setTimeout(() => {
        initSettingsPanel();
    }, 500);
});

// 初始化设置面板
function initSettingsPanel() {
    console.log('初始化设置面板');
    
    // 初始化主题系统
    initThemeSystem();
    
    // 初始化壁纸系统
    initWallpaperSystem();
    
    // 设置标签页切换
    setupSettingsTabs();
    
    // 初始化自动主题功能
    initAutoThemeSettings();
    
    // 初始化搜索功能
    initSettingsSearch();
}

// 设置标签页切换
function setupSettingsTabs() {
    const tabs = document.querySelectorAll('.settings-tab');
    const panels = document.querySelectorAll('.settings-panel');
    
    tabs.forEach(tab => {
        tab.addEventListener('click', () => {
            // 移除所有活动状态
            tabs.forEach(t => t.classList.remove('active'));
            panels.forEach(p => p.classList.remove('active'));
            
            // 添加活动状态
            tab.classList.add('active');
            const panelId = tab.getAttribute('data-tab') + '-panel';
            document.getElementById(panelId).classList.add('active');
        });
    });
}

// 初始化设置搜索
function initSettingsSearch() {
    const searchInput = document.getElementById('settings-search-input');
    if (!searchInput) return;
    
    searchInput.addEventListener('input', () => {
        const searchTerm = searchInput.value.toLowerCase();
        
        // 搜索主题
        const themeOptions = document.querySelectorAll('.theme-option');
        themeOptions.forEach(option => {
            const themeName = option.title.toLowerCase();
            if (searchTerm === '' || themeName.includes(searchTerm)) {
                option.style.display = 'block';
            } else {
                option.style.display = 'none';
            }
        });
        
        // 搜索壁纸
        const wallpaperOptions = document.querySelectorAll('.wallpaper-option');
        wallpaperOptions.forEach(option => {
            const wallpaperName = option.title.toLowerCase();
            if (searchTerm === '' || wallpaperName.includes(searchTerm)) {
                option.style.display = 'block';
            } else {
                option.style.display = 'none';
            }
        });
        
        // 如果有搜索内容，自动显示所有面板
        if (searchTerm !== '') {
            document.querySelectorAll('.settings-panel').forEach(panel => {
                panel.classList.add('active');
            });
        } else {
            // 恢复只显示活动面板
            document.querySelectorAll('.settings-panel').forEach(panel => {
                panel.classList.remove('active');
            });
            const activeTabId = document.querySelector('.settings-tab.active').getAttribute('data-tab') + '-panel';
            document.getElementById(activeTabId).classList.add('active');
        }
    });
}

// 初始化主题系统
function initThemeSystem() {
    console.log('初始化主题系统');
    
    // 加载保存的主题
    loadSavedTheme();
    
    // 设置主题切换事件
    const themeOptions = document.querySelectorAll('.theme-option');
    themeOptions.forEach(option => {
        // 点击事件 - 永久切换主题
        option.addEventListener('click', () => {
            const theme = option.getAttribute('data-theme');
            console.log('主题切换: 点击了主题', theme);
            
            // 更新活动状态
            themeOptions.forEach(opt => {
                opt.classList.remove('active');
                console.log('移除活动状态:', opt.getAttribute('data-theme'));
            });
            option.classList.add('active');
            console.log('添加活动状态:', theme);
            
            // 设置主题
            setTheme(theme);
            console.log('应用主题:', theme);
            
            // 保存主题选择
            saveTheme(theme);
            
            // 添加动画效果
            option.classList.add('previewing');
            setTimeout(() => {
                option.classList.remove('previewing');
            }, 500);
        });
        
        // 鼠标悬停事件 - 临时预览主题
        option.addEventListener('mouseenter', () => {
            // 如果已经是活动主题，不需要预览
            if (option.classList.contains('active')) {
                console.log('鼠标悬停: 已是活动主题，跳过预览');
                return;
            }
            
            // 保存当前主题
            const currentTheme = document.body.getAttribute('data-theme') || 'default';
            option.dataset.previousTheme = currentTheme;
            console.log('鼠标悬停: 保存当前主题', currentTheme);
            
            // 临时应用新主题
            const previewTheme = option.getAttribute('data-theme');
            setTheme(previewTheme, true);
            console.log('鼠标悬停: 预览主题', previewTheme);
            
            // 添加预览指示
            option.classList.add('previewing');
        });
        
        // 鼠标离开事件 - 恢复原主题
        option.addEventListener('mouseleave', () => {
            // 如果是活动主题，不需要恢复
            if (option.classList.contains('active')) {
                console.log('鼠标离开: 已是活动主题，不恢复');
                option.classList.remove('previewing');
                return;
            }
            
            // 如果有保存的主题，恢复它
            if (option.dataset.previousTheme) {
                const previousTheme = option.dataset.previousTheme;
                console.log('鼠标离开: 恢复之前的主题', previousTheme);
                setTheme(previousTheme, true);
                delete option.dataset.previousTheme;
            }
            
            // 移除预览指示
            option.classList.remove('previewing');
        });
    });
    
    // 添加主题切换器的折叠/展开功能
    const themeSwitcher = document.getElementById('theme-switcher');
    if (themeSwitcher) {
        // 获取折叠/展开按钮
        const toggleButton = themeSwitcher.querySelector('.theme-toggle-btn');
        
        if (toggleButton) {
            // 设置按钮标题
            toggleButton.title = getText('theme_settings');
            
            // 默认折叠状态
            let isCollapsed = localStorage.getItem('theme-switcher-collapsed') === 'true';
            
            // 应用初始状态
            if (isCollapsed) {
                themeSwitcher.classList.add('collapsed');
            }
            
            // 点击事件
            toggleButton.addEventListener('click', () => {
                isCollapsed = !isCollapsed;
                if (isCollapsed) {
                    themeSwitcher.classList.add('collapsed');
                } else {
                    themeSwitcher.classList.remove('collapsed');
                }
                localStorage.setItem('theme-switcher-collapsed', isCollapsed);
            });
        }
    }
    
    // 根据系统偏好设置初始主题
    if (!localStorage.getItem('preferred-theme')) {
        // 检测系统偏好
        if (window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches) {
            setTheme('dark');
            document.querySelector('.theme-option[data-theme="dark"]').classList.add('active');
        } else {
            setTheme('default');
            document.querySelector('.theme-option[data-theme="default"]').classList.add('active');
        }
    }
    
    // 监听系统主题变化
    if (window.matchMedia) {
        window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', e => {
            if (!localStorage.getItem('preferred-theme')) {
                setTheme(e.matches ? 'dark' : 'default');
                
                // 更新活动状态
                document.querySelectorAll('.theme-option').forEach(opt => opt.classList.remove('active'));
                document.querySelector(`.theme-option[data-theme="${e.matches ? 'dark' : 'default'}"]`).classList.add('active');
            }
        });
    }
    
    // 更新主题选项标题为当前语言
    updateThemeTitles();
}

// 更新主题选项标题
function updateThemeTitles() {
    document.querySelectorAll('.theme-option').forEach(option => {
        const theme = option.getAttribute('data-theme');
        switch (theme) {
            case 'default':
                option.title = getText('default_theme');
                break;
            case 'dark':
                option.title = getText('dark_theme');
                break;
            case 'teal':
                option.title = getText('teal_theme');
                break;
            case 'purple':
                option.title = getText('purple_theme');
                break;
            case 'orange':
                option.title = getText('orange_theme');
                break;
            case 'pink':
                option.title = getText('pink_theme');
                break;
            case 'green':
                option.title = getText('green_theme');
                break;
            case 'yellow':
                option.title = getText('yellow_theme');
                break;
            case 'red':
                option.title = getText('red_theme');
                break;
            case 'high-contrast':
                option.title = getText('high_contrast');
                break;
        }
    });
    
    // 更新主题切换按钮标题
    const themeToggleBtn = document.querySelector('.theme-toggle-btn');
    if (themeToggleBtn) {
        themeToggleBtn.title = getText('theme_settings');
    }
    
    // 更新设置标签
    const themeTab = document.querySelector('.settings-tab[data-tab="themes"]');
    if (themeTab) {
        themeTab.textContent = getText('themes');
    }
    
    const wallpapersTab = document.querySelector('.settings-tab[data-tab="wallpapers"]');
    if (wallpapersTab) {
        wallpapersTab.textContent = getText('wallpapers');
    }
    
    const autoTab = document.querySelector('.settings-tab[data-tab="auto"]');
    if (autoTab) {
        autoTab.textContent = getText('auto_settings');
    }
    
    // 更新搜索框占位符
    const searchInput = document.getElementById('settings-search-input');
    if (searchInput) {
        searchInput.placeholder = getText('search_settings');
    }
}

// 初始化自动主题功能
function initAutoThemeSettings() {
    // 检查是否启用了自动主题
    const autoThemeEnabled = localStorage.getItem('auto-theme-enabled') === 'true';
    
    // 获取自动主题开关
    const autoThemeToggle = document.getElementById('auto-theme-toggle');
    if (autoThemeToggle) {
        // 设置初始状态
        autoThemeToggle.checked = autoThemeEnabled;
        
        // 设置变更事件
        autoThemeToggle.addEventListener('change', () => {
            localStorage.setItem('auto-theme-enabled', autoThemeToggle.checked);
            
            // 更新UI状态
            const dayThemeSelect = document.querySelector('.day-theme-select');
            const nightThemeSelect = document.querySelector('.night-theme-select');
            
            if (dayThemeSelect && nightThemeSelect) {
                dayThemeSelect.style.display = autoThemeToggle.checked ? 'block' : 'none';
                nightThemeSelect.style.display = autoThemeToggle.checked ? 'block' : 'none';
            }
            
            if (autoThemeToggle.checked) {
                // 启用自动主题
                setTimeBasedTheme();
                // 每小时检查一次
                setInterval(setTimeBasedTheme, 60 * 60 * 1000);
            }
        });
    }
    
    // 填充主题下拉选择框
    populateThemeSelects();
    
    // 设置下拉框初始值
    const dayThemeSelect = document.getElementById('day-theme-select');
    const nightThemeSelect = document.getElementById('night-theme-select');
    
    if (dayThemeSelect && nightThemeSelect) {
        dayThemeSelect.value = localStorage.getItem('day-theme') || 'default';
        nightThemeSelect.value = localStorage.getItem('night-theme') || 'dark';
        
        // 设置变更事件
        dayThemeSelect.addEventListener('change', () => {
            localStorage.setItem('day-theme', dayThemeSelect.value);
            if (autoThemeEnabled) {
                setTimeBasedTheme();
            }
        });
        
        nightThemeSelect.addEventListener('change', () => {
            localStorage.setItem('night-theme', nightThemeSelect.value);
            if (autoThemeEnabled) {
                setTimeBasedTheme();
            }
        });
        
        // 更新UI状态
        dayThemeSelect.parentElement.style.display = autoThemeEnabled ? 'block' : 'none';
        nightThemeSelect.parentElement.style.display = autoThemeEnabled ? 'block' : 'none';
    }
    
    // 如果启用了自动主题，设置基于时间的主题
    if (autoThemeEnabled) {
        setTimeBasedTheme();
        // 每小时检查一次
        setInterval(setTimeBasedTheme, 60 * 60 * 1000);
    }
    
    // 更新文本
    updateAutoThemeText();
}

// 填充主题下拉选择框
function populateThemeSelects() {
    const dayThemeSelect = document.getElementById('day-theme-select');
    const nightThemeSelect = document.getElementById('night-theme-select');
    
    if (!dayThemeSelect || !nightThemeSelect) return;
    
    // 清空现有选项
    dayThemeSelect.innerHTML = '';
    nightThemeSelect.innerHTML = '';
    
    // 添加主题选项
    const themes = [
        { value: 'default', text: getText('default_theme') },
        { value: 'dark', text: getText('dark_theme') },
        { value: 'teal', text: getText('teal_theme') },
        { value: 'purple', text: getText('purple_theme') },
        { value: 'orange', text: getText('orange_theme') },
        { value: 'pink', text: getText('pink_theme') },
        { value: 'green', text: getText('green_theme') },
        { value: 'yellow', text: getText('yellow_theme') },
        { value: 'red', text: getText('red_theme') },
        { value: 'high-contrast', text: getText('high_contrast') }
    ];
    
    themes.forEach(theme => {
        const dayOption = document.createElement('option');
        dayOption.value = theme.value;
        dayOption.textContent = theme.text;
        dayThemeSelect.appendChild(dayOption);
        
        const nightOption = document.createElement('option');
        nightOption.value = theme.value;
        nightOption.textContent = theme.text;
        nightThemeSelect.appendChild(nightOption);
    });
}

// 更新自动主题设置的文本
function updateAutoThemeText() {
    const autoThemeToggle = document.getElementById('auto-theme-toggle');
    if (autoThemeToggle) {
        const toggleLabel = autoThemeToggle.nextElementSibling;
        if (toggleLabel) {
            toggleLabel.textContent = getText('auto_theme');
        }
    }
    
    const dayThemeLabel = document.querySelector('.day-theme-select label');
    if (dayThemeLabel) {
        dayThemeLabel.textContent = getText('day_theme');
    }
    
    const nightThemeLabel = document.querySelector('.night-theme-select label');
    if (nightThemeLabel) {
        nightThemeLabel.textContent = getText('night_theme');
    }
    
    // 更新设置标签
    const autoTab = document.querySelector('.settings-tab[data-tab="auto"]');
    if (autoTab) {
        autoTab.textContent = getText('auto_settings');
    }
}

// 初始化壁纸系统
function initWallpaperSystem() {
    console.log('初始化壁纸系统');
    
    // 加载保存的壁纸
    loadSavedWallpaper();
    
    // 设置壁纸切换事件
    const wallpaperOptions = document.querySelectorAll('.wallpaper-option');
    wallpaperOptions.forEach(option => {
        // 点击事件 - 永久切换壁纸
        option.addEventListener('click', () => {
            const wallpaper = option.getAttribute('data-wallpaper');
            console.log('壁纸切换: 点击了壁纸', wallpaper);
            
            // 更新活动状态
            wallpaperOptions.forEach(opt => {
                opt.classList.remove('active');
                console.log('移除活动状态:', opt.getAttribute('data-wallpaper'));
            });
            option.classList.add('active');
            console.log('添加活动状态:', wallpaper);
            
            // 设置壁纸
            setWallpaper(wallpaper);
            console.log('应用壁纸:', wallpaper);
            
            // 保存壁纸选择
            saveWallpaper(wallpaper);
            
            // 添加动画效果
            option.classList.add('previewing');
            setTimeout(() => {
                option.classList.remove('previewing');
            }, 500);
        });
        
        // 鼠标悬停事件 - 临时预览壁纸
        option.addEventListener('mouseenter', () => {
            // 如果已经是活动壁纸，不需要预览
            if (option.classList.contains('active')) {
                console.log('鼠标悬停: 已是活动壁纸，跳过预览');
                return;
            }
            
            // 保存当前壁纸
            const currentWallpaper = document.body.getAttribute('data-wallpaper') || 'none';
            option.dataset.previousWallpaper = currentWallpaper;
            console.log('鼠标悬停: 保存当前壁纸', currentWallpaper);
            
            // 临时应用新壁纸
            const previewWallpaper = option.getAttribute('data-wallpaper');
            setWallpaper(previewWallpaper, true);
            console.log('鼠标悬停: 预览壁纸', previewWallpaper);
            
            // 添加预览指示
            option.classList.add('previewing');
        });
        
        // 鼠标离开事件 - 恢复原壁纸
        option.addEventListener('mouseleave', () => {
            // 如果是活动壁纸，不需要恢复
            if (option.classList.contains('active')) {
                console.log('鼠标离开: 已是活动壁纸，不恢复');
                option.classList.remove('previewing');
                return;
            }
            
            // 如果有保存的壁纸，恢复它
            if (option.dataset.previousWallpaper) {
                const previousWallpaper = option.dataset.previousWallpaper;
                console.log('鼠标离开: 恢复之前的壁纸', previousWallpaper);
                setWallpaper(previousWallpaper, true);
                delete option.dataset.previousWallpaper;
            }
            
            // 移除预览指示
            option.classList.remove('previewing');
        });
    });
    
    // 更新壁纸选项标题为当前语言
    updateWallpaperTitles();
}

// 设置壁纸
function setWallpaper(wallpaper, isPreview = false) {
    console.log('setWallpaper 函数被调用，壁纸:', wallpaper);
    
    // 移除所有壁纸类
    document.body.classList.remove('bg-pattern-1', 'bg-pattern-2', 'bg-pattern-3', 'bg-pattern-4', 'bg-pattern-5', 
                                  'bg-pattern-6', 'bg-pattern-7', 'bg-pattern-8', 'bg-pattern-9', 'bg-pattern-10');
    
    // 移除壁纸属性
    document.body.removeAttribute('data-wallpaper');
    console.log('移除了现有壁纸');
    
    // 设置新壁纸
    if (wallpaper !== 'none') {
        document.body.setAttribute('data-wallpaper', wallpaper);
        
        // 根据壁纸类型添加对应的类
        const patternNumber = wallpaper.replace('pattern-', '');
        document.body.classList.add('bg-pattern-' + patternNumber);
        
        console.log('设置了新壁纸:', wallpaper);
    } else {
        console.log('使用无壁纸模式');
    }
    
    // 如果不是预览，触发壁纸变更事件
    if (!isPreview) {
        // 创建自定义事件
        const event = new CustomEvent('wallpaperchange', { detail: { wallpaper } });
        document.dispatchEvent(event);
        console.log('触发了wallpaperchange事件');
    }
}

// 保存壁纸选择
function saveWallpaper(wallpaper) {
    // 本地存储
    localStorage.setItem('preferred-wallpaper', wallpaper);
    
    // 保存到后端
    window.go.main.App.SaveWallpaperConfig({wallpaper: wallpaper})
        .then(() => {
            console.log('壁纸配置已保存到后端');
        })
        .catch(error => {
            console.error('保存壁纸配置失败:', error);
            // 如果后端API不存在，只保存到本地
        });
}

// 加载保存的壁纸
function loadSavedWallpaper() {
    console.log('加载保存的壁纸');
    
    // 尝试从后端加载
    if (window.go && window.go.main && window.go.main.App && window.go.main.App.GetWallpaperConfig) {
        window.go.main.App.GetWallpaperConfig()
            .then(config => {
                console.log('从后端加载壁纸配置:', config);
                if (config && config.wallpaper) {
                    setWallpaper(config.wallpaper);
                    
                    // 更新活动状态
                    document.querySelectorAll('.wallpaper-option').forEach(opt => opt.classList.remove('active'));
                    const activeOption = document.querySelector(`.wallpaper-option[data-wallpaper="${config.wallpaper}"]`);
                    if (activeOption) {
                        activeOption.classList.add('active');
                    }
                    
                    // 同步到本地存储
                    localStorage.setItem('preferred-wallpaper', config.wallpaper);
                } else {
                    // 如果后端没有配置，尝试从本地存储加载
                    loadFromLocalStorage();
                }
            })
            .catch(error => {
                console.error('加载壁纸配置失败:', error);
                // 如果后端加载失败，尝试从本地存储加载
                loadFromLocalStorage();
            });
    } else {
        // 如果后端API不存在，直接从本地存储加载
        loadFromLocalStorage();
    }
    
    function loadFromLocalStorage() {
        const savedWallpaper = localStorage.getItem('preferred-wallpaper');
        if (savedWallpaper) {
            setWallpaper(savedWallpaper);
            
            // 更新活动状态
            document.querySelectorAll('.wallpaper-option').forEach(opt => opt.classList.remove('active'));
            const activeOption = document.querySelector(`.wallpaper-option[data-wallpaper="${savedWallpaper}"]`);
            if (activeOption) {
                activeOption.classList.add('active');
            }
            
            // 如果后端API存在，尝试保存到后端
            if (window.go && window.go.main && window.go.main.App && window.go.main.App.SaveWallpaperConfig) {
                window.go.main.App.SaveWallpaperConfig({wallpaper: savedWallpaper}).catch(console.error);
            }
        } else {
            // 默认无壁纸
            setWallpaper('none');
            const noneOption = document.querySelector('.wallpaper-option[data-wallpaper="none"]');
            if (noneOption) {
                noneOption.classList.add('active');
            }
        }
    }
}

// 更新壁纸选项标题
function updateWallpaperTitles() {
    document.querySelectorAll('.wallpaper-option').forEach(option => {
        const wallpaper = option.getAttribute('data-wallpaper');
        switch (wallpaper) {
            case 'none':
                option.title = getText('no_wallpaper');
                break;
            case 'pattern-1':
                option.title = getText('pattern_1');
                break;
            case 'pattern-2':
                option.title = getText('pattern_2');
                break;
            case 'pattern-3':
                option.title = getText('pattern_3');
                break;
            case 'pattern-4':
                option.title = getText('pattern_4');
                break;
            case 'pattern-5':
                option.title = getText('pattern_5');
                break;
            case 'pattern-6':
                option.title = getText('pattern_6');
                break;
            case 'pattern-7':
                option.title = getText('pattern_7');
                break;
            case 'pattern-8':
                option.title = getText('pattern_8');
                break;
            case 'pattern-9':
                option.title = getText('pattern_9');
                break;
            case 'pattern-10':
                option.title = getText('pattern_10');
                break;
        }
    });
    
    // 更新壁纸切换按钮标题
    const wallpaperToggleBtn = document.querySelector('.wallpaper-toggle-btn');
    if (wallpaperToggleBtn) {
        wallpaperToggleBtn.title = getText('wallpaper_settings');
    }
}

// 在语言变化时更新壁纸相关文本
document.addEventListener('languagechange', () => {
    // 更新壁纸选项标题
    updateWallpaperTitles();
}); 