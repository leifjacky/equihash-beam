cmake_minimum_required(VERSION 2.8.4)

set(LIBBEAM_SRC 
	./pow/equihash.cpp 
	./core/ecc.cpp 
	./core/uintBig.cpp 
	./core/difficulty.cpp 
	./core/ecc_bulletproof.cpp 
	./3rdparty/crypto/blake/ref/blake2b-ref.c 
	./3rdparty/crypto/equihash_impl.cpp 
	./3rdparty/utilstrencodings.cpp 
	./3rdparty/arith_uint256.cpp 
	./3rdparty/uint256.cpp
	./utility/common.cpp
	)

include_directories(.)
include_directories(./3rdparty)
include_directories(./3rdparty/secp256k1-zkp/src)

add_library(beam_shared SHARED ${LIBBEAM_SRC})

set_target_properties(beam_shared PROPERTIES OUTPUT_NAME "beampow")

set(LIBRARY_OUTPUT_PATH ${PROJECT_BINARY_DIR}/lib)

install(TARGETS beam_shared
	LIBRARY DESTINATION lib
	ARCHIVE DESTINATION lib)
