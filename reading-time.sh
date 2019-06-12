#!/bin/bash
TIME_PER_WORD=0.34

# TODO: 
# - Check if notification daemon is running.
# - Check if xclip is not empty.

# Get highlighted text from clipboard.
highlighted_text=$(xclip -o)

# Count number of words.
number_of_words=$(echo "$highlighted_text" | wc -w)

# Calculate reading time.
reading_time=$(echo "($number_of_words * $TIME_PER_WORD) / 60" | bc)

if [ $reading_time -le 1 ]; then
    reading_time=1
fi

# Cut off text after 500 characters.
notify-send "Text: '$(echo $highlighted_text | cut -c 1-500)'"
notify-send "Words: $number_of_words words."
notify-send "Reading time: $reading_time minute(s)."

exit 0

