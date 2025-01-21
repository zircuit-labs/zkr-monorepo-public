# This script:
# fetches KZGparams from the s3 bucket. This is needed for the provers.

paramFiles=(
  "k_16.srs"
  "k_17.srs"
  "k_18.srs"
  "k_19.srs"
  "k_20.srs"
  "k_21.srs"
  "k_22.srs"
  "k_23.srs"
  "k_24.srs"
  "k_25.srs"
  "k_26.srs"
)

fetch_url () {
    BASE_URL="https://kparam-bucket-dev.s3.amazonaws.com"
    filename="$@"
    URL="$BASE_URL/$filename"
    target_path="params/$filename"

    # Check if the file already exists
    if [ -e "$target_path" ]; then
      echo "File $filename already exists. Skipping download."
    else
      echo "Fetching ${filename}"
      if curl -L -f -o "$target_path" "$URL"; then
        echo "Downloaded ${filename}"
      else
        echo "Failed to download ${URL}"
        exit 1  # Exit the script with a non-zero status code
      fi
    fi
}

mkdir -p params
for i in "${paramFiles[@]}"
do
  fetch_url "$i"
done
