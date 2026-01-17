/**
 * Truncates the string to comply with the character limit of `maxLength`. If `maxLength` is not
 * provided, the default limit is 50.
 */
export function truncate(str: string, maxLength?: number): string {
    const len: number = maxLength ?? 50;
    if (str.length > len) return str.substring(0, len) + "...";
    return str;
};

/**
 * Displays the time that lapsed from the `date` in seconds, minutes, hours, days or years.
 * @returns A string that indicates how much time have passed from now.
 */
export function showLastUpdated(date: string): string {
    type TimeUnit = "second" | "minute" | "hour" | "day" | "year";

    function format(value: number, unit: TimeUnit): string {
        return `${value} ${unit}${value !== 1 ? "s" : ""} ago`;
    }

    const diffMs: number = Date.now() - new Date(date).valueOf();

    if (diffMs <= 0) return "just now";

    const seconds: number = Math.floor(diffMs / 1000);
    if (seconds < 60) return format(seconds, "second");

    const minutes: number = Math.floor(seconds / 60);
    if (minutes < 60) return format(minutes, "minute");

    const hours: number = Math.floor(minutes / 60);
    if (hours < 24) return format(hours, "hour");

    const days: number = Math.floor(hours / 24);
    if (days < 365) return format(days, "day");

    const years: number = Math.floor(days / 365);
    return format(years, "year");
}
