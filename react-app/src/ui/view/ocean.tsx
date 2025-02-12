import React, { useRef, useEffect } from "react";

const OceanAnimation: React.FC = () => {
    const canvasRef = useRef<HTMLCanvasElement | null>(null);

    useEffect(() => {
        const canvas = canvasRef.current;
        const ctx = canvas?.getContext("2d");
        if (!canvas || !ctx) return;

        const width = canvas.width = window.innerWidth;
        const height = canvas.height = 500;

        const waves = [
            { amplitude: 20, frequency: 0.05, speed: 0.02, offset: 0, direction: 1, level: 0, color: "#00BFFF" },  // Светло-синий
            { amplitude: 30, frequency: 0.04, speed: 0.03, offset: 0, direction: -1, level: 50, color: "#1E90FF" }, // Темно-синий
            { amplitude: 25, frequency: 0.07, speed: 0.05, offset: 0, direction: 1, level: 100, color: "#4682B4" }, // Стальной синий
            { amplitude: 15, frequency: 0.03, speed: 0.015, offset: 0, direction: -1, level: 150, color: "#20B2AA" }, // Темный бирюзовый
            { amplitude: 10, frequency: 0.1, speed: 0.06, offset: 0, direction: 1, level: 200, color: "#5F9EA0" },  // Голубой
        ];

        const drawWaves = () => {
            ctx.clearRect(0, 0, width, height); // Очистить канвас

            waves.forEach((wave) => {
                ctx.beginPath();
                ctx.moveTo(0, height);

                for (let x = 0; x < width; x++) {
                    const y = Math.sin(x * wave.frequency + wave.offset) * wave.amplitude + wave.level + height / 2;
                    ctx.lineTo(x, y);
                }

                ctx.lineTo(width, height);
                ctx.closePath();
                ctx.fillStyle = wave.color; // Устанавливаем цвет волны
                ctx.fill();

                wave.offset += wave.speed * wave.direction; // Двигаем волну для анимации
            });

            requestAnimationFrame(drawWaves);
        };

        drawWaves();

        return () => cancelAnimationFrame(drawWaves as unknown as number); // Очистить анимацию при размонтировании компонента
    }, []);

    return (
        <div style={{ position: "relative", height: "500px" }}>
            <canvas ref={canvasRef}></canvas>
        </div>
    );
};

export default OceanAnimation;
