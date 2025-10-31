const canvas = document.getElementById('gameCanvas');
const ctx = canvas.getContext('2d');
const startBtn = document.getElementById('startBtn');
const scoreElement = document.getElementById('score');

const gridSize = 20;
const tileCount = canvas.width / gridSize;
let snake = [];
let food = {};
let dx = gridSize;
let dy = 0;
let score = 0;
let gameLoop;
let isGameRunning = false;

function initGame() {
    snake = [{x: 5*gridSize, y: 5*gridSize}];
    food = generateFood();
    dx = gridSize;
    dy = 0;
    score = 0;
    scoreElement.textContent = `得分: ${score}`;
}

function generateFood() {
    return {
        x: Math.floor(Math.random() * tileCount) * gridSize,
        y: Math.floor(Math.random() * tileCount) * gridSize
    };
}

function drawGame() {
    // 移动蛇身
    const head = {x: snake[0].x + dx, y: snake[0].y + dy};
    snake.unshift(head);

    // 吃食物检测
    if (head.x === food.x && head.y === food.y) {
        score += 10;
        scoreElement.textContent = `得分: ${score}`;
        food = generateFood();
    } else {
        snake.pop();
    }

    // 清空画布
    ctx.fillStyle = 'white';
    ctx.fillRect(0, 0, canvas.width, canvas.height);

    // 绘制蛇
    ctx.fillStyle = '#4CAF50';
    snake.forEach(segment => {
        ctx.fillRect(segment.x, segment.y, gridSize-2, gridSize-2);
    });

    // 绘制食物
    ctx.fillStyle = 'red';
    ctx.fillRect(food.x, food.y, gridSize-2, gridSize-2);

    // 碰撞检测
    if (head.x < 0 || head.x >= canvas.width || 
        head.y < 0 || head.y >= canvas.height ||
        snake.slice(1).some(segment => segment.x === head.x && segment.y === head.y)) {
        gameOver();
    }
}

function gameOver() {
    clearInterval(gameLoop);
    isGameRunning = false;
    alert(`游戏结束！得分：${score}`);
    startBtn.textContent = '重新开始';
}

startBtn.addEventListener('click', () => {
    if (!isGameRunning) {
        isGameRunning = true;
        startBtn.textContent = '游戏中...';
        initGame();
        gameLoop = setInterval(drawGame, 100);
    }
});

document.addEventListener('keydown', (e) => {
    if (!isGameRunning) return;
    
    switch(e.key) {
        case 'ArrowUp':
            if (dy === 0) { dx = 0; dy = -gridSize; }
            break;
        case 'ArrowDown':
            if (dy === 0) { dx = 0; dy = gridSize; }
            break;
        case 'ArrowLeft':
            if (dx === 0) { dx = -gridSize; dy = 0; }
            break;
        case 'ArrowRight':
            if (dx === 0) { dx = gridSize; dy = 0; }
            break;
    }
});