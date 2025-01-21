import { Logger, createLogger, format, transports } from 'winston';

interface CustomLogger extends Logger {
  done: (msg: string) => void;
}

const { combine, timestamp, printf } = format;

const levels = {
  error: 0,
  warn: 1,
  info: 2,
  http: 3,
  verbose: 4,
  debug: 5,
  silly: 6,
  done: 7,
};

const levelColors: Record<string, string> = {
  error: '31', // Red
  warn: '33', // Yellow
  info: '32', // Green
  http: '35', // Magenta
  verbose: '36', // Cyan
  debug: '34', // Blue
  silly: '37', // White
  done: '34', // Blue for 'done'
};

// Custom printf formatter
const customFormat = printf(({ level, message, timestamp: time }) => {
  const colorCode = levelColors[level] || '37'; // Default to white if no specific color
  // Apply the color to the prefix and include a checkmark for "done" messages
  if (level === 'done') {
    return `${time} \x1b[${colorCode}mDONE:\x1b[0m ${message} ✅`;
  }
  return `${time} \x1b[${colorCode}m${level.toUpperCase()}:\x1b[0m ${message}`;
});

const logLevel = process.env.logDebug === 'true' ? 'done' : 'info';

export const logger = createLogger({
  levels,
  level: logLevel,
  format: combine(
    timestamp({ format: 'HH:mm:ss' }), // Format timestamp to include only hour, minute, and second
    customFormat
  ),
  transports: [new transports.Console()],
}) as CustomLogger;

logger.done = function (msg) {
  this.log({
    level: 'done',
    message: msg,
  });
};

export default logger;
