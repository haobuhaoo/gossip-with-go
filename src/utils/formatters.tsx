/**
 * Truncates the string to comply with the character limit of `maxLength`. If `maxLength` is not
 * provided, the default limit is 50.
 */
export function truncate(str: string, maxLength?: number): string {
    const len = maxLength ?? 50;
    if (str.length > len) {
        return str.substring(0, len) + "...";
    }
    return str;
};