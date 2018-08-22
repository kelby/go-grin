/// Computes the proof-of-work difficulty that the next block should comply
/// with. Takes an iterator over past blocks, from latest (highest height) to
/// oldest (lowest height). The iterator produces pairs of timestamp and
/// difficulty for each block.
///
/// The difficulty calculation is based on both Digishield and GravityWave
/// family of difficulty computation, coming to something very close to Zcash.
/// The reference difficulty is an average of the difficulty over a window of
/// DIFFICULTY_ADJUST_WINDOW blocks. The corresponding timespan is calculated
/// by using the difference between the median timestamps at the beginning
/// and the end of the window.
func Next_difficulty<T>(cursor: T) -> Result<Difficulty, TargetError>
where
  T: IntoIterator<Item = Result<(u64, Difficulty), TargetError>>,
{
  // Create vector of difficulty data running from earliest
  // to latest, and pad with simulated pre-genesis data to allow earlier
  // adjustment if there isn't enough window data
  // length will be DIFFICULTY_ADJUST_WINDOW+MEDIAN_TIME_WINDOW
  let diff_data = global::difficulty_data_to_vector(cursor);

  // Obtain the median window for the earlier time period
  // the first MEDIAN_TIME_WINDOW elements
  let mut window_earliest: Vec<u64> = diff_data
    .iter()
    .take(MEDIAN_TIME_WINDOW as usize)
    .map(|n| n.clone().unwrap().0)
    .collect();
  // pick median
  window_earliest.sort();
  let earliest_ts = window_earliest[MEDIAN_TIME_INDEX as usize];

  // Obtain the median window for the latest time period
  // i.e. the last MEDIAN_TIME_WINDOW elements
  let mut window_latest: Vec<u64> = diff_data
    .iter()
    .skip(DIFFICULTY_ADJUST_WINDOW as usize)
    .map(|n| n.clone().unwrap().0)
    .collect();
  // pick median
  window_latest.sort();
  let latest_ts = window_latest[MEDIAN_TIME_INDEX as usize];

  // median time delta
  let ts_delta = latest_ts - earliest_ts;

  // Get the difficulty sum of the last DIFFICULTY_ADJUST_WINDOW elements
  let diff_sum = diff_data
    .iter()
    .skip(MEDIAN_TIME_WINDOW as usize)
    .fold(0, |sum, d| sum + d.clone().unwrap().1.to_num());

  // Apply dampening except when difficulty is near 1
  let ts_damp = if diff_sum < DAMP_FACTOR * DIFFICULTY_ADJUST_WINDOW {
    ts_delta
  } else {
    (1 * ts_delta + (DAMP_FACTOR - 1) * BLOCK_TIME_WINDOW) / DAMP_FACTOR
  };

  // Apply time bounds
  let adj_ts = if ts_damp < LOWER_TIME_BOUND {
    LOWER_TIME_BOUND
  } else if ts_damp > UPPER_TIME_BOUND {
    UPPER_TIME_BOUND
  } else {
    ts_damp
  };

  let difficulty = diff_sum * BLOCK_TIME_SEC / adj_ts;

  Ok(Difficulty::from_num(max(difficulty, 1)))
}
