import { useTranslation } from "react-i18next"
import { FontAwesomeIcon } from "@fortawesome/react-fontawesome"
import { AnimatePresence, motion } from "framer-motion"
import { TERipple } from "tw-elements-react"

function CopyButton({ text, isVisible }) {
    const { t } = useTranslation()

    const handleCopy = () => {
        navigator.clipboard.writeText(text)
        window.flash("info.successful_copy", "success")
    }
    return (
        <AnimatePresence>
            {isVisible && (
                <motion.div
                    initial={{ opacity: 0 }}
                    animate={{ opacity: 1, transition: { duration: 0.2 } }}
                    exit={{ opacity: 0, transition: { duration: 0.1 } }}>
                    <TERipple className="rounded-lg overflow-hidden" rippleColor="#808080">
                        <button
                            className={`rounded-lg text-grey-200 bg-grey-775 hover:bg-grey-750 border border-borxstrcol relative h-9 w-9`}
                            aria-label={t("aria_label.copy")}
                            onClick={handleCopy}>
                            <FontAwesomeIcon
                                icon="fa-regular fa-copy"
                                className="absolute top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2 w-3.5 h-3.5"
                            />
                        </button>
                    </TERipple>
                </motion.div>
            )}
        </AnimatePresence>
    )
}

export default CopyButton
